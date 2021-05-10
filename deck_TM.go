package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "TM",
		Description: "A custom Trailmakers pack",
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
