package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "reject",
		Description: "Reject Packs - 1, 2, and 3",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},
			&PromptCard{Prompt: `Test Test Test %s Test.`},

		},
		Responses: []ResponseCard{
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`Placeholder`,
			`%blank`,
		},
	}

	AddPack(pack)
}
