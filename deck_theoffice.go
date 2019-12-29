package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "the-office",
		Description: "The Office (US) TV Show Themed",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `The bigger the %s, the more powerful the man. - Dwight Schrute`},
			&PromptCard{Prompt: `What is the most romantic thing in The Office? %s`},
			&PromptCard{Prompt: `Why was The Office so great? %s`},
			&PromptCard{Prompt: `Assistant to the Regional %s.`},
			&PromptCard{Prompt: `Break me off a piece of that %s.`},
			&PromptCard{Prompt: `%s is not a joke Jim! Millions of families suffer every year!`},
			&PromptCard{Prompt: `You don't know me; you've just seen my %s.`},
			&PromptCard{Prompt: `Dwight, At 8AM today, someone poisons %s. Do not %s. More instructions will follow. Cordially, future Dwight.`},
			&PromptCard{Prompt: `%s. Nope. Don't like that.`},
			&PromptCard{Prompt: `I started with a thumb tack and traded my way to a telescope, but in a way the most valuable thing here wasn't the telescope... It was %s.`},
			&PromptCard{Prompt: `The dundie for funniest in the office goes to %s.`},
			&PromptCard{Prompt: `I need a username, and... I have a great one. "%s." That way people will know exactly where my priorities are at.`},
			&PromptCard{Prompt: `Beer me %s.`},
			&PromptCard{Prompt: `Dwight, you %s!`},
			&PromptCard{Prompt: `I have flaws. What are they? I sing in the shower. Sometimes I spend too much time volunteering. Occasionally I'll %s. So sue me.`},
			&PromptCard{Prompt: `Well happy birthday Jesus. Sorry your %s's so lame.`},
			&PromptCard{Prompt: `If I have to do this based on stereotypes that are totally untrue, that I do not agree with, you would maybe not be a very good driver. "Oh, man, am I %s?"`},
			&PromptCard{Prompt: `Friends joke with one another. "Hey, you're poor." "Hey, your momma's dead." "Hey you're %s."`},
			&PromptCard{Prompt: `Hey Mr. Scott, whatcha gonna do, whatcha gonna do? %s.`},
			&PromptCard{Prompt: `The doobie doobie pothead stoner of the year dundie goes to %s.`},
			&PromptCard{Prompt: `What is Michael's worst fear? %s`},
			&PromptCard{Prompt: `What does Michael hate more than Toby? %s`},
			&PromptCard{Prompt: `The subject of today's conference meeting: %s.`},
			&PromptCard{Prompt: `The party planning committee decided on a %s themed Christmas party this year. %s`},
			&PromptCard{Prompt: `Bears. Beets. %s.`},
			&PromptCard{Prompt: `If I had a gun with two bullets and I was in a room with %s and Toby, I would shoot Toby twice.`},
			&PromptCard{Prompt: `%s. That's what she said!`},
			&PromptCard{Prompt: `%s. How the turntables...`},
			&PromptCard{Prompt: `You miss 100% of the %s you don't take. - Wayne Gretzky - Michael Scott`},
			&PromptCard{Prompt: `What's the ratio of Stanley Nickels to Schrute Bucks? The same as the ratio of %s to %s.`},
			&PromptCard{Prompt: `Make %s first, make sales second, make love third. In no particular order.`},
			&PromptCard{Prompt: `You know what's even cooler than a triceratops? %s.`},
			&PromptCard{Prompt: `Jim put my %s in jello again!`},
			&PromptCard{Prompt: `Who should've replaced Michael as regional manager?`},
			&PromptCard{Prompt: `The hardest thing about working at Dunder Mifflin is %s.`},
			&PromptCard{Prompt: `What makes Dwight, Dwight? %s`},
		},
		Responses: []ResponseCard{
			`The length of Roy and Pam's engagement`,
			`Gabe Lewis`,
			`Clark Green`,
			`David Wallace`,
			`Jo Bennett`,
			`Prison Mike`,
			`Date Mike's charm`,
			`Paying $500 for a baggie of Caprese salad`,
			`Kissing a male employee to show the office you aren't homophobic`,
			`Scranton, Pennsylvania`,
			`The Michael Scott Paper Company`,
			`Dunder Mifflin`,
			`Michael's proposal to Holly`,
			`DID I STUTTER?!`,
			`Running over a co-worker with your car`,
			`A well-timed, "That's what she said."`,
			`The love story of Michael and Holly`,
			`The Office after Michael leaves the show`,
			`Exclaiming, "Buttlicker! Our prices have never been lower!"`,
			`Calling someone ASAP as possible`,
			`Presents. The best way to say, "I love you this many dollars-worth."`,
			`Following your GPS into a lake`,
			`Threat Level: Midnight`,
			`The love story of Dwight and Angela`,
			`Going twenty eight years without having sex`,
			`The dundies`,
			`Scott's tots`,
			`The love that Angela has for her cats`,
			`Asian Jim`,
			`Stanley Hudson`,
			`Jan Levinson`,
			`Pam Beesly`,
			`Michael Scott`,
			`Wanting people to be afraid of how much they love you`,
			`Ignorant slut`,
			`A two-way petting zoo`,
			`Not superstitious, but a little stitious`,
			`Dinkin'flicka`,
			`Soup snakes`,
			`Hotel hell. Check-in time is now. Check-out time is never`,
			`Believing in something. Even if it means sacrificing everything`,
			`When your awesome proposal is stolen by a coworker`,
			`Asking your date if they think you're going to have sex so you can plan what you're going to eat`,
			`Rewarding your employee by offering to rent them a room in your house`,
			`Having Kelly's well-balanced love life`,
			`Deangelo's inner circle`,
			`Being desperate enough to seek out Meredith for a good time`,
			`When your black employee is being a debbie-downer so you ask him to go to the back of the bus`,
			`The Battle of Schrute Farms`,
			`When your gay-dar beeps on yourself`,
			`Kevin's toupee`,
			`A 45-day, 45 point plan`,
			`Ryan's sense of business ethics`,
			`Mega desk`,
			`Michael's reaction to Toby's return`,
			`Wedding: the fusing of two metals with a hot torch. - Webster's Dictionary`,
			`The ability to be impregnated. The driving force between male-female attraction`,
			`Proposing to the love your life... in the rain... at a gas station..`,
			`The issue of Philip's paternity`,
			`Letting your boss fall into a koi pond`,
			`Jan's daughter Assturd`,
			`Waking up and stepping on your hot George Foreman grill`,
			`Getting 3 months paid vacation and a company car because your boss outed you to the entire office`,
			`A 60 acre working beet farm`,
			`Competing for gold, silver, and bronze yogurt lids`,
			`Changing Secret Santa to Yankee Swap at the last minute`,
			`Jim's proposal to Pam`,
			`A sign that says "It is your birthday."`,
			`The Scranton Strangler`,
			`Informing every woman you've ever dated that you have herpes`,
			`The love story of Michael and Ryan`,
			`The love story of Jim and Pam`,
			`Professor Copperfield's miracle legumes`,
			`Asking yourself, "Would an idiot do that?"`,
			`Kicking your date in the face at a wedding`,
			`Having the strength of a grown man and a little baby`,
			`Feeling God in this Chili's tonight`,
			`Agent Michael Scarn`,
			`Deangelo Vickers`,
			`Karen Filippelli`,
			`Pete "Plop" Miller`,
			`Holly Flax`,
			`Darryl Philbin`,
			`Kelly Kapoor`,
			`Angela Martin`,
			`Kevin Malone`,
			`Robert California`,
			`Ryan Howard`,
			`Dwight Schrute`,
			`Little Kid Lover`,
			`Blind Guy McSqueezy`,
			`Auctioning off people, like in the olden days`,
			`Football cream`,
			`Going mach five`,
			`That face Jim makes`,
			`DECLARING BANKRUPTCY!`,
			`Kevin's famous chili.`,
			`Meredith's casual friday dress`,
			`Declaring yourself "Hay King" to the shock of the crowd`,
			`Finally fulfilling your lifelong dream of getting promoted to regional manager, accidentally firing a gun at work, and getting demoted`,
			`Charles Miner`,
			`Todd Packer`,
			`Nellie Bertram`,
			`Kelly Erin Hannon`,
			`Toby Flenderson`,
			`Oscar Martinez`,
			`Michael Scott's Dunder-Mifflin Scranton Meredith Palmer Memorial Celebrity Rabies Awareness Pro-Am Fun Run Race For The Cure`,
			`Working fo HR which means not really being a part of the family, also being divorced which means not really being a part of your family`,
			`Jim's cute teapot gift for Pam`,
			`Pam-pong, Angela's favorite office game`,
			`When your husband has an affair... with your male co-worker`,
			`The sinking feeling when Jim declares his love for Pam, and she rejects him`,
			`When you hate to see someone leave, but love to watch them go. Cause of their butt`,
			`The love that the Myrick family has for The Office`,
			`Cutting off a manikin's face and wearing it as a mask`,
			`Signing something that you don't want to sign as "Daffy Duck."`,
			`Hardcore parkour`,
			`Dr. Crentist, the Dentist`,
			`Flonkerton. THe national sport of Icelandic paper companies`,
			`A "World's Best Boss" mug that you bought for yourself`,
			`Mose Schrute`,
			`Perfectenschlag. Defined as when a man's life comes together perfectly, or perfect pork anus`,
			`Saying you were raped to get out of any situation`,
			`Pam's face when Jim asks her out for the first time`,
			`Bob Vance, Vance Refridgeration`,
			`Creed Bratton`,
			`Phyllis Lapin`,
			`Meredith Palmer`,
			`Roy Anderson`,
			`Andy Bernard`,
			`Jim Halpert`,
			`Not fitting in a rowboat`,
			`Paying way too much for worms`,
			`Pretending you're talking to the camera man until the cops leave`,
			`Justice Beaver, the crime-fighting beaver`,
			`Pippity poppity, gimme the zoppity`,
			`Michael Klump`,
			`That scummy microphone guy, Brian`,
			`When a girl that you think likes you is being nice because they think you're retarded`,
			`The Andy Bernard school of anger management`,
			`Pretzel Day`,
			`Beets`,
			`The Sabre Pyramid Tablet`,
			`That's what she said`,
		},
	}

	AddPack(pack)
}
