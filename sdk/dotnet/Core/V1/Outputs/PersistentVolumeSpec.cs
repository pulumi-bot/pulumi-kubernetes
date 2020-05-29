// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    [OutputType]
    public sealed class PersistentVolumeSpec
    {
        /// <summary>
        /// AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
        /// </summary>
        public readonly ImmutableArray<string> AccessModes;
        /// <summary>
        /// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.AWSElasticBlockStoreVolumeSource AwsElasticBlockStore;
        /// <summary>
        /// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.AzureDiskVolumeSource AzureDisk;
        /// <summary>
        /// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.AzureFilePersistentVolumeSource AzureFile;
        /// <summary>
        /// A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
        /// </summary>
        public readonly ImmutableDictionary<string, string> Capacity;
        /// <summary>
        /// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.CephFSPersistentVolumeSource Cephfs;
        /// <summary>
        /// Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.CinderPersistentVolumeSource Cinder;
        /// <summary>
        /// ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference ClaimRef;
        /// <summary>
        /// CSI represents storage that is handled by an external CSI driver (Beta feature).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.CSIPersistentVolumeSource Csi;
        /// <summary>
        /// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.FCVolumeSource Fc;
        /// <summary>
        /// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.FlexPersistentVolumeSource FlexVolume;
        /// <summary>
        /// Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.FlockerVolumeSource Flocker;
        /// <summary>
        /// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.GCEPersistentDiskVolumeSource GcePersistentDisk;
        /// <summary>
        /// Glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.GlusterfsPersistentVolumeSource Glusterfs;
        /// <summary>
        /// HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.HostPathVolumeSource HostPath;
        /// <summary>
        /// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ISCSIPersistentVolumeSource Iscsi;
        /// <summary>
        /// Local represents directly-attached storage with node affinity
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.LocalVolumeSource Local;
        /// <summary>
        /// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
        /// </summary>
        public readonly ImmutableArray<string> MountOptions;
        /// <summary>
        /// NFS represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NFSVolumeSource Nfs;
        /// <summary>
        /// NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.VolumeNodeAffinity NodeAffinity;
        /// <summary>
        /// What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
        /// </summary>
        public readonly string PersistentVolumeReclaimPolicy;
        /// <summary>
        /// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PhotonPersistentDiskVolumeSource PhotonPersistentDisk;
        /// <summary>
        /// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PortworxVolumeSource PortworxVolume;
        /// <summary>
        /// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.QuobyteVolumeSource Quobyte;
        /// <summary>
        /// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.RBDPersistentVolumeSource Rbd;
        /// <summary>
        /// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ScaleIOPersistentVolumeSource ScaleIO;
        /// <summary>
        /// Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
        /// </summary>
        public readonly string StorageClassName;
        /// <summary>
        /// StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.StorageOSPersistentVolumeSource Storageos;
        /// <summary>
        /// volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
        /// </summary>
        public readonly string VolumeMode;
        /// <summary>
        /// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.VsphereVirtualDiskVolumeSource VsphereVolume;

        [OutputConstructor]
        private PersistentVolumeSpec(
            ImmutableArray<string> accessModes,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.AWSElasticBlockStoreVolumeSource awsElasticBlockStore,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.AzureDiskVolumeSource azureDisk,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.AzureFilePersistentVolumeSource azureFile,

            ImmutableDictionary<string, string> capacity,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.CephFSPersistentVolumeSource cephfs,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.CinderPersistentVolumeSource cinder,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference claimRef,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.CSIPersistentVolumeSource csi,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.FCVolumeSource fc,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.FlexPersistentVolumeSource flexVolume,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.FlockerVolumeSource flocker,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.GCEPersistentDiskVolumeSource gcePersistentDisk,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.GlusterfsPersistentVolumeSource glusterfs,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.HostPathVolumeSource hostPath,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ISCSIPersistentVolumeSource iscsi,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.LocalVolumeSource local,

            ImmutableArray<string> mountOptions,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.NFSVolumeSource nfs,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.VolumeNodeAffinity nodeAffinity,

            string persistentVolumeReclaimPolicy,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PhotonPersistentDiskVolumeSource photonPersistentDisk,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PortworxVolumeSource portworxVolume,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.QuobyteVolumeSource quobyte,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.RBDPersistentVolumeSource rbd,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ScaleIOPersistentVolumeSource scaleIO,

            string storageClassName,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.StorageOSPersistentVolumeSource storageos,

            string volumeMode,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.VsphereVirtualDiskVolumeSource vsphereVolume)
        {
            AccessModes = accessModes;
            AwsElasticBlockStore = awsElasticBlockStore;
            AzureDisk = azureDisk;
            AzureFile = azureFile;
            Capacity = capacity;
            Cephfs = cephfs;
            Cinder = cinder;
            ClaimRef = claimRef;
            Csi = csi;
            Fc = fc;
            FlexVolume = flexVolume;
            Flocker = flocker;
            GcePersistentDisk = gcePersistentDisk;
            Glusterfs = glusterfs;
            HostPath = hostPath;
            Iscsi = iscsi;
            Local = local;
            MountOptions = mountOptions;
            Nfs = nfs;
            NodeAffinity = nodeAffinity;
            PersistentVolumeReclaimPolicy = persistentVolumeReclaimPolicy;
            PhotonPersistentDisk = photonPersistentDisk;
            PortworxVolume = portworxVolume;
            Quobyte = quobyte;
            Rbd = rbd;
            ScaleIO = scaleIO;
            StorageClassName = storageClassName;
            Storageos = storageos;
            VolumeMode = volumeMode;
            VsphereVolume = vsphereVolume;
        }
    }
}
