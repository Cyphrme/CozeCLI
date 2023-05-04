package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cristalhq/acmd"
	"github.com/cyphrme/coze"
)

func main() {

	cmds := []acmd.Command{
		{
			Name:        "sign",
			Description: "Sign coze with given key. (coze sign <coze> <key>)",
			ExecFunc: func(ctx context.Context, args []string) error {
				cz := new(coze.Coze)
				err := json.Unmarshal([]byte(args[0]), cz)
				if err != nil {
					panic(err)
				}

				key := new(coze.Key)
				err = json.Unmarshal([]byte(args[1]), key)
				if err != nil {
					panic(err)
				}

				err = key.SignCoze(cz)
				if err != nil {
					panic(err)
				}
				fmt.Println(cz)

				return nil
			},
		},

		{
			Name:        "signpay",
			Description: "Sign pay with given key. (coze signpay <pay> <key>)",
			ExecFunc: func(ctx context.Context, args []string) error {
				pay := new(coze.Pay)
				err := json.Unmarshal([]byte(args[0]), pay)
				if err != nil {
					panic(err)
				}

				key := new(coze.Key)
				err = json.Unmarshal([]byte(args[1]), key)
				if err != nil {
					panic(err)
				}

				cz, err := key.SignPay(pay)
				if err != nil {
					panic(err)
				}
				fmt.Println(cz)

				return nil
			},
		},

		{
			Name:        "verify",
			Description: "Verify coze with given key. (coze verify <coze> <key>)",
			ExecFunc: func(ctx context.Context, args []string) error {
				cz := new(coze.Coze)
				err := json.Unmarshal([]byte(args[0]), cz)
				if err != nil {
					panic(err)
				}

				key := new(coze.Key)
				err = json.Unmarshal([]byte(args[1]), key)
				if err != nil {
					panic(err)
				}

				b, err := key.VerifyCoze(cz)
				if err != nil {
					panic(err)
				}
				fmt.Println(b)

				return nil
			},
		},

		{
			Name:        "newkey",
			Description: "Generate a new key. (coze newkey [<alg>])",
			ExecFunc: func(ctx context.Context, args []string) error {
				var alg coze.SEAlg
				if len(args) == 0 {
					alg = coze.SEAlg(coze.ES256) // TODO change default to Ed25519 once JS coze supports Ed25519
				} else {
					alg = coze.SEAlg(coze.Parse(args[0]))
				}

				key, err := coze.NewKey(alg)
				if err != nil {
					panic(err)
				}
				fmt.Println(key)
				return nil
			},
		},

		{
			Name:        "tmb",
			Description: "Calculate tmb for coze key. (coze tmb <key>)",
			ExecFunc: func(ctx context.Context, args []string) error {
				key := new(coze.Key)
				err := json.Unmarshal([]byte(args[0]), key)
				if err != nil {
					panic(err)
				}

				key.Thumbprint()
				if err != nil {
					panic(err)
				}
				fmt.Println(key.Tmb)

				return nil
			},
		},

		{
			Name:        "meta",
			Description: "Calculate meta for coze. (coze meta <coze>)",
			ExecFunc: func(ctx context.Context, args []string) error {
				cz := new(coze.Coze)
				err := json.Unmarshal([]byte(args[0]), cz)
				if err != nil {
					panic(err)
				}

				err = cz.Meta()
				if err != nil {
					panic(err)
				}

				cz.Pay = nil
				cz.Sig = nil

				fmt.Println(cz)
				return nil
			},
		},

		{
			Name:        "revoke",
			Description: "Generate revoke coze for private key.  (coze revoke <key>)",
			ExecFunc: func(ctx context.Context, args []string) error {
				key := new(coze.Key)
				err := json.Unmarshal([]byte(args[0]), key)
				if err != nil {
					panic(err)
				}

				cz, err := key.Revoke()
				if err != nil {
					panic(err)
				}
				fmt.Println(cz)
				return nil
			},
		},
	} // End of cmds

	// all the acmd.Config fields are optional
	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName:        "Coze CLI",
		AppDescription: "CLI for Coze, a cryptographic JSON messaging specification designed for human readability.",
		Version:        "v0.0.0",
		// Context - if nil `signal.Notify` will be used
		// Args - if nil `os.Args[1:]` will be used
		// Usage - if nil default print will be used
	})

	if err := r.Run(); err != nil {
		r.Exit(err)
	}

}

type flags struct {
	// key json.RawMessage

}
