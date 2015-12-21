package fv

// the references: in order: of the Base set of Fighter Verses
var BaseSet = []entry{
	{ BibleOrder: 4, SetOrder: 101, Set: "Base", Ref: "Deuteronomy 7:9", },
	{ BibleOrder: 5, SetOrder: 102, Set: "Base", Ref: "Deuteronomy 10:12–13", },
	{ BibleOrder: 133, SetOrder: 103, Set: "Base", Ref: "John 1:12–13", },
	{ BibleOrder: 164, SetOrder: 104, Set: "Base", Ref: "Romans 11:33–36", },
	{ BibleOrder: 165, SetOrder: 105, Set: "Base", Ref: "Romans 12:1-2", },
	{ BibleOrder: 38, SetOrder: 106, Set: "Base", Ref: "Psalm 56:3-4", },
	{ BibleOrder: 39, SetOrder: 107, Set: "Base", Ref: "Psalm 62:5-7 [8]", },
	{ BibleOrder: 155, SetOrder: 108, Set: "Base", Ref: "Romans 8:1", },
	{ BibleOrder: 253, SetOrder: 109, Set: "Base", Ref: "1 John 2:15-17", },
	{ BibleOrder: 207, SetOrder: 110, Set: "Base", Ref: "Philippians 2:5-7", },
	{ BibleOrder: 208, SetOrder: 111, Set: "Base", Ref: "Philippians 2:8-9", },
	{ BibleOrder: 209, SetOrder: 112, Set: "Base", Ref: "Philippians 2:10-11", },
	{ BibleOrder: 210, SetOrder: 113, Set: "Base", Ref: "Philippians 2:12-13", },
	{ BibleOrder: 236, SetOrder: 114, Set: "Base", Ref: "James 1:2-3", },
	{ BibleOrder: 237, SetOrder: 115, Set: "Base", Ref: "James 1:4-5", },
	{ BibleOrder: 8, SetOrder: 116, Set: "Base", Ref: "Psalm 1:1-2", },
	{ BibleOrder: 9, SetOrder: 117, Set: "Base", Ref: "Psalm 1:3-4", },
	{ BibleOrder: 10, SetOrder: 118, Set: "Base", Ref: "Psalm 1:5-6", },
	{ BibleOrder: 218, SetOrder: 119, Set: "Base", Ref: "Colossians 3:1-3", },
	{ BibleOrder: 198, SetOrder: 120, Set: "Base", Ref: "Ephesians 4:26", },
	{ BibleOrder: 104, SetOrder: 121, Set: "Base", Ref: "Isaiah 40:28-29", },
	{ BibleOrder: 105, SetOrder: 122, Set: "Base", Ref: "Isaiah 40:30-31", },
	{ BibleOrder: 44, SetOrder: 123, Set: "Base", Ref: "Psalm 86:5-7", },
	{ BibleOrder: 224, SetOrder: 124, Set: "Base", Ref: "1 Timothy 2:5", },
	{ BibleOrder: 245, SetOrder: 125, Set: "Base", Ref: "1 Peter 1:3–5", },
	{ BibleOrder: 201, SetOrder: 126, Set: "Base", Ref: "Ephesians 6:10-11", },
	{ BibleOrder: 202, SetOrder: 127, Set: "Base", Ref: "Ephesians 6:12-13", },
	{ BibleOrder: 203, SetOrder: 128, Set: "Base", Ref: "Ephesians 6:14-15", },
	{ BibleOrder: 204, SetOrder: 129, Set: "Base", Ref: "Ephesians 6:16-17 [18]", },
	{ BibleOrder: 205, SetOrder: 130, Set: "Base", Ref: "Philippians 1:6", },
	{ BibleOrder: 127, SetOrder: 131, Set: "Base", Ref: "Matthew 10:28", },
	{ BibleOrder: 149, SetOrder: 132, Set: "Base", Ref: "Romans 1:16 [17]", },
	{ BibleOrder: 128, SetOrder: 133, Set: "Base", Ref: "Matthew 11:28-30", },
	{ BibleOrder: 16, SetOrder: 134, Set: "Base", Ref: "Psalm 20:6–7 [8]", },
	{ BibleOrder: 238, SetOrder: 135, Set: "Base", Ref: "James 1:12", },
	{ BibleOrder: 185, SetOrder: 136, Set: "Base", Ref: "2 Corinthians 9:6-7", },
	{ BibleOrder: 186, SetOrder: 137, Set: "Base", Ref: "2 Corinthians 9:8", },
	{ BibleOrder: 187, SetOrder: 138, Set: "Base", Ref: "2 Corinthians 12:9 [10]", },
	{ BibleOrder: 113, SetOrder: 139, Set: "Base", Ref: "Isaiah 64:4", },
	{ BibleOrder: 228, SetOrder: 140, Set: "Base", Ref: "Titus 3:4-6", },
	{ BibleOrder: 110, SetOrder: 141, Set: "Base", Ref: "Isaiah 46:9-10 [11]", },
	{ BibleOrder: 85, SetOrder: 142, Set: "Base", Ref: "Proverbs 1:10", },
	{ BibleOrder: 86, SetOrder: 143, Set: "Base", Ref: "Proverbs 3:5-6 [7]", },
	{ BibleOrder: 97, SetOrder: 144, Set: "Base", Ref: "Proverbs 19:11", },
	{ BibleOrder: 144, SetOrder: 145, Set: "Base", Ref: "John 15:5", },
	{ BibleOrder: 142, SetOrder: 146, Set: "Base", Ref: "John 14:2–3", },
	{ BibleOrder: 73, SetOrder: 147, Set: "Base", Ref: "Psalm 125:1-2", },
	{ BibleOrder: 84, SetOrder: 148, Set: "Base", Ref: "Psalm 141:3-4", },
	{ BibleOrder: 252, SetOrder: 149, Set: "Base", Ref: "1 John 1:8-9", },
	{ BibleOrder: 17, SetOrder: 150, Set: "Base", Ref: "Psalm 23:1-2", },
	{ BibleOrder: 18, SetOrder: 151, Set: "Base", Ref: "Psalm 23:3-4", },
	{ BibleOrder: 19, SetOrder: 152, Set: "Base", Ref: "Psalm 23:5-6", },
	{ BibleOrder: 103, SetOrder: 201, Set: "Base", Ref: "Isaiah 40:8", },
	{ BibleOrder: 162, SetOrder: 202, Set: "Base", Ref: "Romans 10:13-14 [15]", },
	{ BibleOrder: 12, SetOrder: 203, Set: "Base", Ref: "Psalm 16:11", },
	{ BibleOrder: 172, SetOrder: 204, Set: "Base", Ref: "Romans 15:1-2", },
	{ BibleOrder: 59, SetOrder: 205, Set: "Base", Ref: "Psalm 103:1-4", },
	{ BibleOrder: 60, SetOrder: 206, Set: "Base", Ref: "Psalm 103:5-7", },
	{ BibleOrder: 61, SetOrder: 207, Set: "Base", Ref: "Psalm 103:8-10", },
	{ BibleOrder: 62, SetOrder: 208, Set: "Base", Ref: "Psalm 103:11-14", },
	{ BibleOrder: 63, SetOrder: 209, Set: "Base", Ref: "Psalm 103:15-16", },
	{ BibleOrder: 64, SetOrder: 210, Set: "Base", Ref: "Psalm 103:17–19", },
	{ BibleOrder: 65, SetOrder: 211, Set: "Base", Ref: "Psalm 103:20-22", },
	{ BibleOrder: 45, SetOrder: 212, Set: "Base", Ref: "Psalm 86:11", },
	{ BibleOrder: 199, SetOrder: 213, Set: "Base", Ref: "Ephesians 4:29", },
	{ BibleOrder: 200, SetOrder: 214, Set: "Base", Ref: "Ephesians 4:31–32", },
	{ BibleOrder: 2, SetOrder: 215, Set: "Base", Ref: "Deuteronomy 6:4–5", },
	{ BibleOrder: 3, SetOrder: 216, Set: "Base", Ref: "Deuteronomy 6:6–7", },
	{ BibleOrder: 182, SetOrder: 217, Set: "Base", Ref: "2 Corinthians 5:17", },
	{ BibleOrder: 131, SetOrder: 218, Set: "Base", Ref: "Luke 12:32–34", },
	{ BibleOrder: 189, SetOrder: 219, Set: "Base", Ref: "Galatians 5:22-23", },
	{ BibleOrder: 190, SetOrder: 220, Set: "Base", Ref: "Galatians 5:24-25", },
	{ BibleOrder: 91, SetOrder: 221, Set: "Base", Ref: "Proverbs 6:20–21", },
	{ BibleOrder: 92, SetOrder: 222, Set: "Base", Ref: "Proverbs 6:22–23", },
	{ BibleOrder: 216, SetOrder: 223, Set: "Base", Ref: "Philippians 4:11-13", },
	{ BibleOrder: 227, SetOrder: 224, Set: "Base", Ref: "2 Timothy 1:7", },
	{ BibleOrder: 250, SetOrder: 225, Set: "Base", Ref: "1 Peter 5:6-8", },
	{ BibleOrder: 251, SetOrder: 226, Set: "Base", Ref: "1 Peter 5:9-10 [11]", },
	{ BibleOrder: 96, SetOrder: 227, Set: "Base", Ref: "Proverbs 18:10", },
	{ BibleOrder: 46, SetOrder: 228, Set: "Base", Ref: "Psalm 91:1-2", },
	{ BibleOrder: 47, SetOrder: 229, Set: "Base", Ref: "Psalm 91:3–4", },
	{ BibleOrder: 48, SetOrder: 230, Set: "Base", Ref: "Psalm 91:5–6", },
	{ BibleOrder: 49, SetOrder: 231, Set: "Base", Ref: "Psalm 91:7-8", },
	{ BibleOrder: 50, SetOrder: 232, Set: "Base", Ref: "Psalm 91:9-10", },
	{ BibleOrder: 51, SetOrder: 233, Set: "Base", Ref: "Psalm 91:11-13", },
	{ BibleOrder: 52, SetOrder: 234, Set: "Base", Ref: "Psalm 91:14-16", },
	{ BibleOrder: 249, SetOrder: 235, Set: "Base", Ref: "1 Peter 4:16", },
	{ BibleOrder: 134, SetOrder: 236, Set: "Base", Ref: "John 3:16–17", },
	{ BibleOrder: 146, SetOrder: 237, Set: "Base", Ref: "Acts 4:11-12", },
	{ BibleOrder: 100, SetOrder: 238, Set: "Base", Ref: "Proverbs 29:1: 11", },
	{ BibleOrder: 217, SetOrder: 239, Set: "Base", Ref: "Philippians 4:19", },
	{ BibleOrder: 175, SetOrder: 240, Set: "Base", Ref: "1 Corinthians 10:13", },
	{ BibleOrder: 111, SetOrder: 241, Set: "Base", Ref: "Isaiah 53:4-5", },
	{ BibleOrder: 112, SetOrder: 242, Set: "Base", Ref: "Isaiah 53:6", },
	{ BibleOrder: 247, SetOrder: 243, Set: "Base", Ref: "1 Peter 2:24", },
	{ BibleOrder: 181, SetOrder: 244, Set: "Base", Ref: "2 Corinthians 4:17-18", },
	{ BibleOrder: 188, SetOrder: 245, Set: "Base", Ref: "Galatians 2:20", },
	{ BibleOrder: 150, SetOrder: 246, Set: "Base", Ref: "Romans 3:23–24", },
	{ BibleOrder: 232, SetOrder: 247, Set: "Base", Ref: "Hebrews 11:6", },
	{ BibleOrder: 171, SetOrder: 248, Set: "Base", Ref: "Romans 14:7-8: [9]", },
	{ BibleOrder: 135, SetOrder: 249, Set: "Base", Ref: "John 3:36", },
	{ BibleOrder: 225, SetOrder: 250, Set: "Base", Ref: "1 Timothy 4:12", },
	{ BibleOrder: 174, SetOrder: 251, Set: "Base", Ref: "1 Corinthians 2:1-2", },
	{ BibleOrder: 257, SetOrder: 252, Set: "Base", Ref: "Revelation 5:12-13", },
	{ BibleOrder: 6, SetOrder: 301, Set: "Base", Ref: "Joshua 1:9", },
	{ BibleOrder: 7, SetOrder: 302, Set: "Base", Ref: "2 Chronicles 16:9", },
	{ BibleOrder: 211, SetOrder: 303, Set: "Base", Ref: "Philippians 3:7-8", },
	{ BibleOrder: 212, SetOrder: 304, Set: "Base", Ref: "Philippians 3:9", },
	{ BibleOrder: 213, SetOrder: 305, Set: "Base", Ref: "Philippians 3:10-11", },
	{ BibleOrder: 173, SetOrder: 306, Set: "Base", Ref: "1 Corinthians 1:18", },
	{ BibleOrder: 33, SetOrder: 307, Set: "Base", Ref: "Psalm 37:[1-2] 3-4", },
	{ BibleOrder: 34, SetOrder: 308, Set: "Base", Ref: "Psalm 37:5-6 [7-8]", },
	{ BibleOrder: 35, SetOrder: 309, Set: "Base", Ref: "Psalm 37:23-24", },
	{ BibleOrder: 233, SetOrder: 310, Set: "Base", Ref: "Hebrews 12:1", },
	{ BibleOrder: 234, SetOrder: 311, Set: "Base", Ref: "Hebrews 12:2", },
	{ BibleOrder: 53, SetOrder: 312, Set: "Base", Ref: "Psalm 96:1-3", },
	{ BibleOrder: 54, SetOrder: 313, Set: "Base", Ref: "Psalm 96:4-5", },
	{ BibleOrder: 55, SetOrder: 314, Set: "Base", Ref: "Psalm 96:6-8", },
	{ BibleOrder: 56, SetOrder: 315, Set: "Base", Ref: "Psalm 96:9-10", },
	{ BibleOrder: 108, SetOrder: 316, Set: "Base", Ref: "Isaiah 43:25", },
	{ BibleOrder: 166, SetOrder: 317, Set: "Base", Ref: "Romans 12:9-10", },
	{ BibleOrder: 167, SetOrder: 318, Set: "Base", Ref: "Romans 12:11-13", },
	{ BibleOrder: 168, SetOrder: 319, Set: "Base", Ref: "Romans 12:14-16", },
	{ BibleOrder: 169, SetOrder: 320, Set: "Base", Ref: "Romans 12:17–19", },
	{ BibleOrder: 170, SetOrder: 321, Set: "Base", Ref: "Romans 12:20-21", },
	{ BibleOrder: 93, SetOrder: 322, Set: "Base", Ref: "Proverbs 15:1", },
	{ BibleOrder: 243, SetOrder: 323, Set: "Base", Ref: "James 4:13-14", },
	{ BibleOrder: 244, SetOrder: 324, Set: "Base", Ref: "James 4:15-17", },
	{ BibleOrder: 132, SetOrder: 325, Set: "Base", Ref: "Luke 19:10", },
	{ BibleOrder: 13, SetOrder: 326, Set: "Base", Ref: "Psalm 18:30-31", },
	{ BibleOrder: 214, SetOrder: 327, Set: "Base", Ref: "Philippians 4:6–7", },
	{ BibleOrder: 215, SetOrder: 328, Set: "Base", Ref: "Philippians 4:8", },
	{ BibleOrder: 36, SetOrder: 329, Set: "Base", Ref: "Psalm 42:11", },
	{ BibleOrder: 109, SetOrder: 330, Set: "Base", Ref: "Isaiah 46:3-4", },
	{ BibleOrder: 206, SetOrder: 331, Set: "Base", Ref: "Philippians 1:21", },
	{ BibleOrder: 116, SetOrder: 332, Set: "Base", Ref: "Jeremiah 29:11-14", },
	{ BibleOrder: 98, SetOrder: 333, Set: "Base", Ref: "Proverbs 22:1", },
	{ BibleOrder: 23, SetOrder: 334, Set: "Base", Ref: "Psalm 30:4-5", },
	{ BibleOrder: 148, SetOrder: 335, Set: "Base", Ref: "Acts 20:35", },
	{ BibleOrder: 122, SetOrder: 336, Set: "Base", Ref: "Matthew 5:3–6", },
	{ BibleOrder: 123, SetOrder: 337, Set: "Base", Ref: "Matthew 5:7–10", },
	{ BibleOrder: 124, SetOrder: 338, Set: "Base", Ref: "Matthew 5:11–12", },
	{ BibleOrder: 177, SetOrder: 339, Set: "Base", Ref: "1 Corinthians 13:4-7", },
	{ BibleOrder: 24, SetOrder: 340, Set: "Base", Ref: "Psalm 32:8 [9]", },
	{ BibleOrder: 101, SetOrder: 341, Set: "Base", Ref: "Proverbs 31:30", },
	{ BibleOrder: 125, SetOrder: 342, Set: "Base", Ref: "Matthew 6:19–21", },
	{ BibleOrder: 176, SetOrder: 343, Set: "Base", Ref: "1 Corinthians 10:31", },
	{ BibleOrder: 153, SetOrder: 344, Set: "Base", Ref: "Romans 5:18-19", },
	{ BibleOrder: 136, SetOrder: 345, Set: "Base", Ref: "John 5:39-40", },
	{ BibleOrder: 246, SetOrder: 346, Set: "Base", Ref: "1 Peter 2:9-10 [11]", },
	{ BibleOrder: 163, SetOrder: 347, Set: "Base", Ref: "Romans 10:17", },
	{ BibleOrder: 129, SetOrder: 348, Set: "Base", Ref: "Matthew 20:26-28", },
	{ BibleOrder: 183, SetOrder: 349, Set: "Base", Ref: "2 Corinthians 5:21", },
	{ BibleOrder: 254, SetOrder: 350, Set: "Base", Ref: "1 John 3:1 [2]", },
	{ BibleOrder: 197, SetOrder: 351, Set: "Base", Ref: "Ephesians 3:20-21", },
	{ BibleOrder: 130, SetOrder: 352, Set: "Base", Ref: "Matthew 22:37-39", },
	{ BibleOrder: 106, SetOrder: 401, Set: "Base", Ref: "Isaiah 41:10", },
	{ BibleOrder: 107, SetOrder: 402, Set: "Base", Ref: "Isaiah 43:1–3", },
	{ BibleOrder: 143, SetOrder: 403, Set: "Base", Ref: "John 14:6", },
	{ BibleOrder: 221, SetOrder: 404, Set: "Base", Ref: "1 Thessalonians 5:14-17", },
	{ BibleOrder: 222, SetOrder: 405, Set: "Base", Ref: "1 Thessalonians 5:18-22", },
	{ BibleOrder: 87, SetOrder: 406, Set: "Base", Ref: "Proverbs 3:9-10", },
	{ BibleOrder: 88, SetOrder: 407, Set: "Base", Ref: "Proverbs 3:11-12", },
	{ BibleOrder: 140, SetOrder: 408, Set: "Base", Ref: "John 10:27-30", },
	{ BibleOrder: 147, SetOrder: 409, Set: "Base", Ref: "Acts 5:29", },
	{ BibleOrder: 57, SetOrder: 410, Set: "Base", Ref: "Psalm 100:1-3", },
	{ BibleOrder: 58, SetOrder: 411, Set: "Base", Ref: "Psalm 100:4-5", },
	{ BibleOrder: 193, SetOrder: 412, Set: "Base", Ref: "Ephesians 2:1-3", },
	{ BibleOrder: 194, SetOrder: 413, Set: "Base", Ref: "Ephesians 2:4-5", },
	{ BibleOrder: 195, SetOrder: 414, Set: "Base", Ref: "Ephesians 2:6-7", },
	{ BibleOrder: 196, SetOrder: 415, Set: "Base", Ref: "Ephesians 2:8-10", },
	{ BibleOrder: 151, SetOrder: 416, Set: "Base", Ref: "Romans 5:8", },
	{ BibleOrder: 152, SetOrder: 417, Set: "Base", Ref: "Romans 5:9-10", },
	{ BibleOrder: 75, SetOrder: 418, Set: "Base", Ref: "Psalm 139:1-3", },
	{ BibleOrder: 76, SetOrder: 419, Set: "Base", Ref: "Psalm 139:4-5", },
	{ BibleOrder: 77, SetOrder: 420, Set: "Base", Ref: "Psalm 139:6-8", },
	{ BibleOrder: 78, SetOrder: 421, Set: "Base", Ref: "Psalm 139:9-10", },
	{ BibleOrder: 79, SetOrder: 422, Set: "Base", Ref: "Psalm 139:11-12", },
	{ BibleOrder: 80, SetOrder: 423, Set: "Base", Ref: "Psalm 139:13-14", },
	{ BibleOrder: 81, SetOrder: 424, Set: "Base", Ref: "Psalm 139:15-16", },
	{ BibleOrder: 82, SetOrder: 425, Set: "Base", Ref: "Psalm 139:17-18", },
	{ BibleOrder: 83, SetOrder: 426, Set: "Base", Ref: "Psalm 139:23-24", },
	{ BibleOrder: 95, SetOrder: 427, Set: "Base", Ref: "Proverbs 17:9: 22", },
	{ BibleOrder: 229, SetOrder: 428, Set: "Base", Ref: "Hebrews 1:1-2", },
	{ BibleOrder: 230, SetOrder: 429, Set: "Base", Ref: "Hebrews 1:3-4", },
	{ BibleOrder: 114, SetOrder: 430, Set: "Base", Ref: "Jeremiah 1:12", },
	{ BibleOrder: 11, SetOrder: 431, Set: "Base", Ref: "Psalm 9:9-10", },
	{ BibleOrder: 239, SetOrder: 432, Set: "Base", Ref: "James 1:17", },
	{ BibleOrder: 20, SetOrder: 433, Set: "Base", Ref: "Psalm 27:1 [2-3]", },
	{ BibleOrder: 21, SetOrder: 434, Set: "Base", Ref: "Psalm 27:4 [5]", },
	{ BibleOrder: 22, SetOrder: 435, Set: "Base", Ref: "Psalm 27:13-14", },
	{ BibleOrder: 241, SetOrder: 436, Set: "Base", Ref: "James 1:22-24", },
	{ BibleOrder: 256, SetOrder: 437, Set: "Base", Ref: "Revelation 2:10", },
	{ BibleOrder: 99, SetOrder: 438, Set: "Base", Ref: "Proverbs 26:20", },
	{ BibleOrder: 115, SetOrder: 439, Set: "Base", Ref: "Jeremiah 9:23-24", },
	{ BibleOrder: 154, SetOrder: 440, Set: "Base", Ref: "Romans 6:23", },
	{ BibleOrder: 178, SetOrder: 441, Set: "Base", Ref: "1 Corinthians 15:1-3", },
	{ BibleOrder: 248, SetOrder: 442, Set: "Base", Ref: "1 Peter 3:18", },
	{ BibleOrder: 37, SetOrder: 443, Set: "Base", Ref: "Psalm 55:22", },
	{ BibleOrder: 74, SetOrder: 444, Set: "Base", Ref: "Psalm 127:1", },
	{ BibleOrder: 242, SetOrder: 445, Set: "Base", Ref: "James 4:7-8", },
	{ BibleOrder: 68, SetOrder: 446, Set: "Base", Ref: "Psalm 119:14-16", },
	{ BibleOrder: 145, SetOrder: 447, Set: "Base", Ref: "John 15:7", },
	{ BibleOrder: 67, SetOrder: 448, Set: "Base", Ref: "Psalm 118:13-14", },
	{ BibleOrder: 94, SetOrder: 449, Set: "Base", Ref: "Proverbs 16:32", },
	{ BibleOrder: 231, SetOrder: 450, Set: "Base", Ref: "Hebrews 3:12-13", },
	{ BibleOrder: 180, SetOrder: 451, Set: "Base", Ref: "1 Corinthians 15:58", },
	{ BibleOrder: 141, SetOrder: 452, Set: "Base", Ref: "John 11:25-26", },
	{ BibleOrder: 121, SetOrder: 501, Set: "Base", Ref: "Daniel 2:20-21", },
	{ BibleOrder: 126, SetOrder: 502, Set: "Base", Ref: "Matthew 9:13", },
	{ BibleOrder: 69, SetOrder: 503, Set: "Base", Ref: "Psalm 121:1-2", },
	{ BibleOrder: 70, SetOrder: 504, Set: "Base", Ref: "Psalm 121:3-4", },
	{ BibleOrder: 71, SetOrder: 505, Set: "Base", Ref: "Psalm 121:5-6", },
	{ BibleOrder: 72, SetOrder: 506, Set: "Base", Ref: "Psalm 121:7-8", },
	{ BibleOrder: 156, SetOrder: 507, Set: "Base", Ref: "Romans 8:28", },
	{ BibleOrder: 157, SetOrder: 508, Set: "Base", Ref: "Romans 8:29-30", },
	{ BibleOrder: 158, SetOrder: 509, Set: "Base", Ref: "Romans 8:31-32", },
	{ BibleOrder: 159, SetOrder: 510, Set: "Base", Ref: "Romans 8:33-34", },
	{ BibleOrder: 160, SetOrder: 511, Set: "Base", Ref: "Romans 8:35-37", },
	{ BibleOrder: 161, SetOrder: 512, Set: "Base", Ref: "Romans 8:38-39", },
	{ BibleOrder: 192, SetOrder: 513, Set: "Base", Ref: "Galatians 6:14", },
	{ BibleOrder: 1, SetOrder: 514, Set: "Base", Ref: "Numbers 23:19", },
	{ BibleOrder: 223, SetOrder: 515, Set: "Base", Ref: "1 Timothy 1:15", },
	{ BibleOrder: 235, SetOrder: 516, Set: "Base", Ref: "Hebrews 13:5-6", },
	{ BibleOrder: 118, SetOrder: 517, Set: "Base", Ref: "Lamentations 3:21-23", },
	{ BibleOrder: 119, SetOrder: 518, Set: "Base", Ref: "Lamentations 3:24-26", },
	{ BibleOrder: 120, SetOrder: 519, Set: "Base", Ref: "Lamentations 3:31-33", },
	{ BibleOrder: 219, SetOrder: 520, Set: "Base", Ref: "Colossians 3:16-17", },
	{ BibleOrder: 102, SetOrder: 521, Set: "Base", Ref: "Isaiah 26:3-4", },
	{ BibleOrder: 14, SetOrder: 522, Set: "Base", Ref: "Psalm 19:7-8", },
	{ BibleOrder: 15, SetOrder: 523, Set: "Base", Ref: "Psalm 19:9-11", },
	{ BibleOrder: 137, SetOrder: 524, Set: "Base", Ref: "John 6:35", },
	{ BibleOrder: 191, SetOrder: 525, Set: "Base", Ref: "Galatians 6:9-10", },
	{ BibleOrder: 25, SetOrder: 526, Set: "Base", Ref: "Psalm 34:1-3", },
	{ BibleOrder: 26, SetOrder: 527, Set: "Base", Ref: "Psalm 34:4-5", },
	{ BibleOrder: 27, SetOrder: 528, Set: "Base", Ref: "Psalm 34:6-8", },
	{ BibleOrder: 28, SetOrder: 529, Set: "Base", Ref: "Psalm 34:9-11", },
	{ BibleOrder: 29, SetOrder: 530, Set: "Base", Ref: "Psalm 34:12-14", },
	{ BibleOrder: 30, SetOrder: 531, Set: "Base", Ref: "Psalm 34:15-16", },
	{ BibleOrder: 31, SetOrder: 532, Set: "Base", Ref: "Psalm 34:17-18", },
	{ BibleOrder: 32, SetOrder: 533, Set: "Base", Ref: "Psalm 34:19-22", },
	{ BibleOrder: 220, SetOrder: 534, Set: "Base", Ref: "Colossians 4:6", },
	{ BibleOrder: 138, SetOrder: 535, Set: "Base", Ref: "John 8:31-32", },
	{ BibleOrder: 139, SetOrder: 536, Set: "Base", Ref: "John 10:10", },
	{ BibleOrder: 117, SetOrder: 537, Set: "Base", Ref: "Jeremiah 32:40", },
	{ BibleOrder: 40, SetOrder: 538, Set: "Base", Ref: "Psalm 73:25-26", },
	{ BibleOrder: 89, SetOrder: 539, Set: "Base", Ref: "Proverbs 4:23-24", },
	{ BibleOrder: 90, SetOrder: 540, Set: "Base", Ref: "Proverbs 4:25-27", },
	{ BibleOrder: 240, SetOrder: 541, Set: "Base", Ref: "James 1:19-20", },
	{ BibleOrder: 184, SetOrder: 542, Set: "Base", Ref: "2 Corinthians 8:9", },
	{ BibleOrder: 41, SetOrder: 543, Set: "Base", Ref: "Psalm 77:13-14", },
	{ BibleOrder: 66, SetOrder: 544, Set: "Base", Ref: "Psalm 118:5-8", },
	{ BibleOrder: 226, SetOrder: 545, Set: "Base", Ref: "1 Timothy 6:6-7", },
	{ BibleOrder: 42, SetOrder: 546, Set: "Base", Ref: "Psalm 79:9", },
	{ BibleOrder: 43, SetOrder: 547, Set: "Base", Ref: "Psalm 84:10-11 [12]", },
	{ BibleOrder: 255, SetOrder: 548, Set: "Base", Ref: "1 John 4:4", },
	{ BibleOrder: 179, SetOrder: 549, Set: "Base", Ref: "1 Corinthians 15:51-52", },
	{ BibleOrder: 258, SetOrder: 550, Set: "Base", Ref: "Revelation 21:3", },
	{ BibleOrder: 259, SetOrder: 551, Set: "Base", Ref: "Revelation 21:4", },
	{ BibleOrder: 260, SetOrder: 552, Set: "Base", Ref: "Revelation 21:5-6 [7]", },
}
