// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
)

func GenerateExampleDaemonSetYaml(dsName, namespace string) *appsv1.DaemonSet {
	Expect(dsName).NotTo(BeEmpty())
	Expect(namespace).NotTo(BeEmpty())

	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      dsName,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": dsName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": dsName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "samplepod",
							Image:           "alpine",
							ImagePullPolicy: "IfNotPresent",
							Command:         []string{"/bin/ash", "-c", "sleep infinity"},
						},
					},
				},
			},
		},
	}
}

func CreateSpiderDoctorAgentDaemonSet(name, namespace string, labels, annotations map[string]string) *appsv1.DaemonSet {
	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: annotations,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "spiderdoctor-agent",
							Image:   "ghcr.io/spidernet-io/spiderdoctor-agent:v0.3.0",
							Command: []string{"/usr/bin/agent"},
							Args:    []string{"--config-path=/tmp/config-map/conf.yml"},
							Env: []corev1.EnvVar{
								{
									Name:  "ENV_LOG_LEVEL",
									Value: "info",
								}, {
									Name:  "ENV_ENABLED_METRIC",
									Value: "false",
								}, {
									Name:  "ENV_METRIC_HTTP_PORT",
									Value: "5711",
								}, {
									Name:  "ENV_HTTP_PORT",
									Value: "80",
								}, {
									Name:  "ENV_GOPS_LISTEN_PORT",
									Value: "5712",
								}, {
									Name:  "ENV_AGENT_GRPC_LISTEN_PORT",
									Value: "3000",
								}, {
									Name:  "ENV_CLUSTER_DNS_DOMAIN",
									Value: "cluster.local",
								}, {
									Name:  "ENV_ENABLE_AGGREGATE_AGENT_REPORT",
									Value: "true",
								}, {
									Name:  "ENV_AGENT_REPORT_STORAGE_PATH",
									Value: "/report",
								}, {
									Name:  "ENV_CLEAN_AGED_REPORT_INTERVAL_IN_MINUTE",
									Value: "10",
								}, {
									Name: "ENV_POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.name",
										},
									},
								}, {
									Name: "ENV_LOCAL_NODE_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "status.hostIP",
										},
									},
								}, {
									Name: "ENV_LOCAL_NODE_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "spec.nodeName",
										},
									},
								}, {
									Name: "ENV_POD_NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.namespace",
										},
									},
								},
							},
							LivenessProbe: &corev1.Probe{
								FailureThreshold: 6,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthy/liveness",
										Port: intstr.IntOrString{
											Type:   intstr.Int,
											IntVal: 80,
										},
										Scheme: "http",
									},
								},
								InitialDelaySeconds: 60,
								PeriodSeconds:       10,
								SuccessThreshold:    1,
								TimeoutSeconds:      5,
							},
							ReadinessProbe: &corev1.Probe{
								FailureThreshold: 3,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthy/readness",
										Port: intstr.IntOrString{
											Type:   intstr.Int,
											IntVal: 80,
										},
										Scheme: "http",
									},
								},
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								SuccessThreshold:    1,
								TimeoutSeconds:      5,
							},
							StartupProbe: &corev1.Probe{
								FailureThreshold: 60,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthy/startup",
										Port: intstr.IntOrString{
											Type:   intstr.Int,
											IntVal: 80,
										},
										Scheme: "http",
									},
								},
								PeriodSeconds:    2,
								SuccessThreshold: 1,
								TimeoutSeconds:   1,
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "config-path",
									MountPath: "/tmp/config-map",
									ReadOnly:  true,
								}, {
									Name:      "report-data",
									MountPath: "/report",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"kubernetes.io/os": "linux",
					},
					Volumes: []corev1.Volume{
						{
							Name: "config-path",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									DefaultMode: pointer.Int32(256),
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "spiderdoctor",
									},
								},
							},
						}, {
							Name: "report-data",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/var/run/spiderdoctor/agent",
									Type: (*corev1.HostPathType)(pointer.String(string(corev1.HostPathDirectoryOrCreate))),
								},
							},
						},
					},
				},
			},
		},
	}
}
