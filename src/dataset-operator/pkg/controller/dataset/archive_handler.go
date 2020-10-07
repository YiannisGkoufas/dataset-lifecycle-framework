package dataset

import (
	comv1alpha1 "github.com/IBM/dataset-lifecycle-framework/src/dataset-operator/pkg/apis/com/v1alpha1"
	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	batchv1 "k8s.io/api/batch/v1"
	"path"
)

func getPodDataDownload(dataset *comv1alpha1.Dataset, operatorNamespace string) *batchv1.Job {
	uuid_forpod, _ := uuid.NewUUID()
	podId := uuid_forpod.String()
	fileUrl := dataset.Spec.Url
	fileName := path.Base(fileUrl)
	seconds := int32(60)
	podSpec := corev1.PodSpec{
		Volumes: []corev1.Volume{
			{Name: "minio-pvc", VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: "minio-pvc",
				},
			}},
		},
		Containers: []corev1.Container{{
			Image: "busybox",
			Name:  "busybox",
			Command: []string{
				"/bin/sh", "-c", "mkdir /data/" + podId + " && wget " + fileUrl + " -P" + " /tmp && " + "tar " + "-xf " + "/tmp/" + fileName + " -C /data/" + podId + " && rm -rf /tmp/" + fileName,
			},
			VolumeMounts: []corev1.VolumeMount{
				{Name: "minio-pvc", MountPath: "/data"},
			},
		}},
		RestartPolicy: corev1.RestartPolicyNever,
	}
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podId,
			Namespace: operatorNamespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: podSpec,
			},
			TTLSecondsAfterFinished: &seconds,
		},
	}
	return job
}
