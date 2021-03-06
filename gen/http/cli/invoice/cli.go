// Code generated by goa v3.7.3, DO NOT EDIT.
//
// invoice HTTP client CLI support package
//
// Command:
// $ goa gen github.com/samverrall/spacecms-api/spacecms-api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	authc "github.com/samverrall/spacecms-api/gen/http/auth/client"
	cmsc "github.com/samverrall/spacecms-api/gen/http/cms/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `auth (create-account|grant-token)
cms (create-page|create-template)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auth create-account --body '{
      "email": "Consequatur aut eos est dolor.",
      "id": "Ea alias minus possimus aut.",
      "name": "Ea minus cupiditate.",
      "password": "Omnis ex."
   }'` + "\n" +
		os.Args[0] + ` cms create-page --body '{
      "createdAt": "2014-06-29T21:24:29Z",
      "id": "FF1D889F-0741-6290-783B-66E606310D86",
      "isActive": false,
      "meta": {
         "description": "Excepturi non occaecati odit voluptates deleniti.",
         "title": "Dolores optio."
      },
      "templateId": "Nisi odio velit ducimus facilis molestiae.",
      "url": "Ratione reprehenderit quae voluptas."
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		authFlags = flag.NewFlagSet("auth", flag.ContinueOnError)

		authCreateAccountFlags    = flag.NewFlagSet("create-account", flag.ExitOnError)
		authCreateAccountBodyFlag = authCreateAccountFlags.String("body", "REQUIRED", "")

		authGrantTokenFlags     = flag.NewFlagSet("grant-token", flag.ExitOnError)
		authGrantTokenBodyFlag  = authGrantTokenFlags.String("body", "REQUIRED", "")
		authGrantTokenTokenFlag = authGrantTokenFlags.String("token", "", "")

		cmsFlags = flag.NewFlagSet("cms", flag.ContinueOnError)

		cmsCreatePageFlags     = flag.NewFlagSet("create-page", flag.ExitOnError)
		cmsCreatePageBodyFlag  = cmsCreatePageFlags.String("body", "REQUIRED", "")
		cmsCreatePageTokenFlag = cmsCreatePageFlags.String("token", "", "")

		cmsCreateTemplateFlags     = flag.NewFlagSet("create-template", flag.ExitOnError)
		cmsCreateTemplateBodyFlag  = cmsCreateTemplateFlags.String("body", "REQUIRED", "")
		cmsCreateTemplateTokenFlag = cmsCreateTemplateFlags.String("token", "", "")
	)
	authFlags.Usage = authUsage
	authCreateAccountFlags.Usage = authCreateAccountUsage
	authGrantTokenFlags.Usage = authGrantTokenUsage

	cmsFlags.Usage = cmsUsage
	cmsCreatePageFlags.Usage = cmsCreatePageUsage
	cmsCreateTemplateFlags.Usage = cmsCreateTemplateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "auth":
			svcf = authFlags
		case "cms":
			svcf = cmsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "auth":
			switch epn {
			case "create-account":
				epf = authCreateAccountFlags

			case "grant-token":
				epf = authGrantTokenFlags

			}

		case "cms":
			switch epn {
			case "create-page":
				epf = cmsCreatePageFlags

			case "create-template":
				epf = cmsCreateTemplateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "auth":
			c := authc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-account":
				endpoint = c.CreateAccount()
				data, err = authc.BuildCreateAccountPayload(*authCreateAccountBodyFlag)
			case "grant-token":
				endpoint = c.GrantToken()
				data, err = authc.BuildGrantTokenPayload(*authGrantTokenBodyFlag, *authGrantTokenTokenFlag)
			}
		case "cms":
			c := cmsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-page":
				endpoint = c.CreatePage()
				data, err = cmsc.BuildCreatePagePayload(*cmsCreatePageBodyFlag, *cmsCreatePageTokenFlag)
			case "create-template":
				endpoint = c.CreateTemplate()
				data, err = cmsc.BuildCreateTemplatePayload(*cmsCreateTemplateBodyFlag, *cmsCreateTemplateTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// authUsage displays the usage of the auth command and its subcommands.
func authUsage() {
	fmt.Fprintf(os.Stderr, `RESTFUL API Backend for Invoicify. An open source invoicing web app.
Usage:
    %[1]s [globalflags] auth COMMAND [flags]

COMMAND:
    create-account: Create an account by email address and password.
    grant-token: Create an account by email address and password.

Additional help:
    %[1]s auth COMMAND --help
`, os.Args[0])
}
func authCreateAccountUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth create-account -body JSON

Create an account by email address and password.
    -body JSON: 

Example:
    %[1]s auth create-account --body '{
      "email": "Consequatur aut eos est dolor.",
      "id": "Ea alias minus possimus aut.",
      "name": "Ea minus cupiditate.",
      "password": "Omnis ex."
   }'
`, os.Args[0])
}

func authGrantTokenUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth grant-token -body JSON -token STRING

Create an account by email address and password.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s auth grant-token --body '{
      "email": "Quia illo voluptatem.",
      "password": "Quas quas autem."
   }' --token "Suscipit laboriosam laborum voluptas at."
`, os.Args[0])
}

// cmsUsage displays the usage of the cms command and its subcommands.
func cmsUsage() {
	fmt.Fprintf(os.Stderr, `CMS service for SpaceCMS. An open source content management system.
Usage:
    %[1]s [globalflags] cms COMMAND [flags]

COMMAND:
    create-page: Create a page
    create-template: Create a template.

Additional help:
    %[1]s cms COMMAND --help
`, os.Args[0])
}
func cmsCreatePageUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] cms create-page -body JSON -token STRING

Create a page
    -body JSON: 
    -token STRING: 

Example:
    %[1]s cms create-page --body '{
      "createdAt": "2014-06-29T21:24:29Z",
      "id": "FF1D889F-0741-6290-783B-66E606310D86",
      "isActive": false,
      "meta": {
         "description": "Excepturi non occaecati odit voluptates deleniti.",
         "title": "Dolores optio."
      },
      "templateId": "Nisi odio velit ducimus facilis molestiae.",
      "url": "Ratione reprehenderit quae voluptas."
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func cmsCreateTemplateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] cms create-template -body JSON -token STRING

Create a template.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s cms create-template --body '{
      "blockId": "Animi et iure fugit vitae totam.",
      "createdAt": "1975-01-21T23:59:55Z",
      "id": "74C53540-E974-ABFF-2565-6BF99F9017B2",
      "name": "Atque sint ea laborum."
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}
