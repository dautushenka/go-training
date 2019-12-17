package fixtures

import "go-training/proj/core/model"

var Posts = []model.Post{
	// News
	{
		Id:         1,
		CategoryId: 1,
		Title:      "Pope lifts 'pontifical secret' rule over sex abuse",
		Body:       "Pope Francis lifts the secrecy rule that critics said protected some abusers from investigation.",
	},
	{
		Id:         2,
		CategoryId: 1,
		Title:      "Pakistan ex-leader Musharraf sentenced to death",
		Body:       "Former military chief Pervez Musharraf is convicted in absentia of treason in a long-running trial.",
	},
	{
		Id:         3,
		CategoryId: 1,
		Title:      "Johnson's Brexit bill to block any delay",
		Body:       "The move comes as MPs gather for the first sitting of the new Parliament after the general election.",
	},
	{
		Id:         4,
		CategoryId: 1,
		Title:      "A bumpy ride for Democrat backing impeachment",
		Body:       "Some politicians are facing anger from constituents over the looming impeachment of President Trump.",
	},
	{
		Id:         5,
		CategoryId: 1,
		Title:      "Charlize Theron opens up about her mum killing her dad",
		Body:       "The actress says she's \"not ashamed\" to discuss the violence she experienced in her family.",
	},

	// Sport
	{
		Id:         6,
		CategoryId: 2,
		Title:      "Italian clubs condemn Serie A's anti-racism monkey posters",
		Body:       "AC Milan say they \"strongly disagree\" with, and were not consulted about, the use of monkeys in a Serie A anti-racism campaign.",
	},
	{
		Id:         7,
		CategoryId: 2,
		Title:      "Arsenal target Arteta to travel with Man City for cup tie",
		Body:       "Mikel Arteta will travel with Manchester City for Wednesday's Carabao Cup quarter-final at Oxford despite holding talks with Arsenal about their managerial vacancy, says Pep Guardiola.",
	},
	{
		Id:         8,
		CategoryId: 2,
		Title:      "Carlo Ancelotti is Everton's first choice to be new manager",
		Body:       "Carlo Ancelotti is Everton's first choice to be their new manager with the Toffees set to intensify their attempts to secure a deal for the Italian.",
	},
	{
		Id:         9,
		CategoryId: 2,
		Title:      "English football's moment of the decade: You decide",
		Body:       "As 2019 comes to a close, BBC Sport has gone through the archive from the last 10 years to find a moment of the decade in English football.",
	},
	{
		Id:         10,
		CategoryId: 2,
		Title:      "Anthony Joshua: Tyson Fury or Deontay Wilder bout still potentially two years away",
		Body:       "Boxing politics look set to delay any fight between Anthony Joshua and either Deontay Wilder or Tyson Fury until 2022, according to the 5 Live Boxing team.",
	},

	// WorkLife
	{
		Id:         11,
		CategoryId: 3,
		Title:      "Why is Sweden deporting talented tech workers?",
		Body:       "Retaining highly-skilled foreign workers is essential for countries like Sweden. But as a deported worker sues the state, the tech-savvy nation’s rejection record is in the spotlight.",
	},
	{
		Id:         12,
		CategoryId: 3,
		Title:      "The vegan condoms keeping sex lives safe – and sustainable",
		Body:       "Can your sex life be more sustainable? These German entrepreneurs think so – and have grown a vegan condom brand into a multimillion-euro business to prove it.",
	},
	{
		Id:         13,
		CategoryId: 3,
		Title:      "Should I delete Tinder? These millennials think so",
		Body:       "More than half a decade since dating apps went mainstream, can millennials who’ve lost patience with digital platforms still find love in the analogue world?",
	},
	{
		Id:         14,
		CategoryId: 3,
		Title:      "How to conquer work paralysis like Ernest Hemingway",
		Body:       "The author wasn’t all about literary masterpieces, dry martinis and rakish charm – he also invented a technique that can beat procrastination and boost productivity.",
	},

	//Travel
	{
		Id:         15,
		CategoryId: 4,
		Title:      "The UK's real-life Treasure Island",
		Body:       "Legend has it that author Robert Louis Stevenson’s 1869 visit to the Scottish island inspired his classic tale of adventure, Treasure Island.",
	},
	{
		Id:         16,
		CategoryId: 4,
		Title:      "Italy's village of 'ugly' people",
		Body:       "Celebrating “ugliness” for the past 140 years, Piobbico has become renowned for being the world capital of ugly people. Now, its utopian idea has blossomed into a worldwide movement.",
	},
	{
		Id:         17,
		CategoryId: 4,
		Title:      "Why these people are so good with money",
		Body:       "The Memon predisposition towards frugality is iconic within Pakistan, but they celebrate their stereotyping as an achievement; a tribute to their enduring prosperity and resilience.",
	},
}
