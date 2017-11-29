package main

// passage represents the return payload for v3 of the ESV api
type passage struct {
	Query       string   `json:"query"`
	Canonical   string   `json:"canonical"`
	Parsed      [][]int  `json:"parsed"`
	Passages    []string `json:"passages"`
	PassageMeta []struct {
		Canonical    string `json:"canonical"`
		PrevVerse    int    `json:"prev_verse"`
		NextVerse    int    `json:"next_verse"`
		ChapterStart []int  `json:"chapter_start"`
		ChapterEnd   []int  `json:"chapter_end"`
		PrevChapter  []int  `json:"prev_chapter"`
		NextChapter  []int  `json:"next_chapter"`
	}
}

// {
// 	"query": "Isa 26:3",
// 	"canonical": "Isaiah 26:3",
// 	"parsed": [
// 	  [
// 		23026003,
// 		23026003
// 	  ]
// 	],
// 	"passage_meta": [
// 	  {
// 		"canonical": "Isaiah 26:3",
// 		"chapter_start": [
// 		  23026001,
// 		  23026021
// 		],
// 		"chapter_end": [
// 		  23026001,
// 		  23026021
// 		],
// 		"prev_verse": 23026002,
// 		"next_verse": 23026004,
// 		"prev_chapter": [
// 		  23025001,
// 		  23025012
// 		],
// 		"next_chapter": [
// 		  23027001,
// 		  23027013
// 		]
// 	  }
// 	],
// 	"passages": [
// 	  "\nIsaiah 26:3\n\n    [3] You keep him in perfect peace\n        whose mind is stayed on you,\n        because he trusts in you. (ESV)"
// 	]
//   }
