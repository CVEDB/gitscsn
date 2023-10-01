module github.com/aquasecurity/gitscan/examples/misconf/go-testing

go 1.16

require (
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.29
	github.com/Azure/go-autorest/autorest/adal v0.9.23
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.12
	github.com/CycloneDX/cyclonedx-go v0.7.2
	github.com/GoogleCloudPlatform/docker-credential-gcr v2.0.5+incompatible
	github.com/Masterminds/sprig/v3 v3.2.3
	github.com/NYTimes/gziphandler v1.1.1
	github.com/alicebob/miniredis/v2 v2.30.5
	github.com/aquasecurity/bolt-fixtures v0.0.0-20200903104109-d34e7f983986
	github.com/aquasecurity/defsec v0.93.0
	github.com/aquasecurity/go-dep-parser v0.0.0-20230926074641-48d70d534559
	github.com/aquasecurity/go-gem-version v0.0.0-20201115065557-8eed6fe000ce
	github.com/aquasecurity/go-npm-version v0.0.0-20201110091526-0b796d180798
	github.com/aquasecurity/go-pep440-version v0.0.0-20210121094942-22b2f8951d46
	github.com/aquasecurity/go-version v0.0.0-20210121072130-637058cfe492
	github.com/aquasecurity/loading v0.0.5
	github.com/aquasecurity/memoryfs v1.4.4
	github.com/aquasecurity/table v1.8.0
	github.com/aquasecurity/testdocker v0.0.0-20230706091143-09ed655568da
	github.com/aquasecurity/tml v0.6.1
	github.com/aquasecurity/trivy-db v0.0.0-20230927082224-b7e0a5886daa
	github.com/aws/aws-sdk-go v1.45.19
	github.com/aws/aws-sdk-go-v2 v1.21.0
	github.com/aws/aws-sdk-go-v2/config v1.18.42
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.122.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.22.0
	github.com/caarlos0/env/v6 v6.10.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cheggaaa/pb/v3 v3.1.4
	github.com/containerd/containerd v1.7.6
	github.com/docker/docker v24.0.6+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/fatih/color v1.15.0
	github.com/go-git/go-git/v5 v5.9.0
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang/protobuf v1.5.3
	github.com/google/go-containerregistry v0.16.1
	github.com/google/licenseclassifier/v2 v2.0.0
	github.com/google/uuid v1.3.1
	github.com/google/wire v0.5.0
	github.com/hashicorp/go-getter v1.7.2
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/in-toto/in-toto-golang v0.9.0
	github.com/knqyf263/go-apk-version v0.0.0-20200609155635-041fdbb8563f
	github.com/knqyf263/go-deb-version v0.0.0-20230223133812-3ed183d23422
	github.com/knqyf263/go-rpm-version v0.0.0-20220614171824-631e686d1075
	github.com/knqyf263/go-rpmdb v0.0.0-20230912071815-bd1c2e66bbe9
	github.com/knqyf263/nested v0.0.1
	github.com/kylelemons/godebug v1.1.0
	github.com/mailru/easyjson v0.7.7
	github.com/masahiro331/go-disk v0.0.0-20220919035250-c8da316f91ac
	github.com/masahiro331/go-ebs-file v0.0.0-20230228042409-005c81d4ae43
	github.com/masahiro331/go-ext4-filesystem v0.0.0-20230705164539-c4f6a70cf8be
	github.com/masahiro331/go-mvn-version v0.0.0-20210429150710-d3157d602a08
	github.com/masahiro331/go-vmdk-parser v0.0.0-20221225061455-612096e4bbbd
	github.com/masahiro331/go-xfs-filesystem v0.0.0-20230608043311-a335f4599b70
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/moby/buildkit v0.12.2
	github.com/open-policy-agent/opa v0.57.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.1.0-rc4
	github.com/owenrumney/go-sarif/v2 v2.2.2
	github.com/package-url/packageurl-go v0.1.1
	github.com/samber/lo v1.38.1
	github.com/saracen/walker v0.1.3
	github.com/secure-systems-lab/go-securesystemslib v0.7.0
	github.com/sigstore/rekor v1.3.0
	github.com/sosedoff/gitkit v0.4.0
	github.com/spdx/tools-golang v0.5.3
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	github.com/testcontainers/testcontainers-go v0.25.0
	github.com/tetratelabs/wazero v1.5.0
	github.com/twitchtv/twirp v8.1.3+incompatible
	github.com/xlab/treeprint v1.2.0
	go.etcd.io/bbolt v1.3.7
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.13.0
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
	golang.org/x/mod v0.12.0
	golang.org/x/sync v0.3.0
	golang.org/x/term v0.12.0
	golang.org/x/text v0.13.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
	gotest.tools v2.2.0+incompatible
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b
	modernc.org/sqlite v1.26.0
)
