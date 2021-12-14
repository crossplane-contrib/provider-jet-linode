module github.com/crossplane-contrib/provider-jet-linode

go 1.16

require (
	github.com/aws/aws-sdk-go v1.41.3 // indirect
	github.com/crossplane-contrib/provider-jet-template v0.0.0-20211203145717-1fcc0a115550
	github.com/crossplane-contrib/terrajet v0.2.1
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.9.0
	github.com/hashicorp/terraform-provider-hashicups v0.3.1
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)

// Comment out the line below instead of the above, if your Terraform provider
// uses an old version (<v2) of github.com/hashicorp/terraform-plugin-sdk.
replace github.com/hashicorp/terraform-plugin-sdk => github.com/turkenh/terraform-plugin-sdk v1.17.2-patch1
