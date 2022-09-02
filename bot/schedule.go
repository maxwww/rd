package bot

type SchoolBell struct {
	Position int
	Start    string
	End      string
}

type Day struct {
	Name     string
	Subjects [][]Subject
}

type Subject struct {
	Name      string
	ShortName string
}

var LiteraturneChytannia = Subject{
	Name:      "Літературне читання",
	ShortName: "Літ. читання",
}

var Matematyka = Subject{
	Name:      "Математика",
	ShortName: "Математика",
}

var UkrainskaMova = Subject{
	Name:      "Українська мова",
	ShortName: "Укр. мова",
}

var Informatyka = Subject{
	Name:      "Інформатика",
	ShortName: "Інформатика",
}

var MuzychneMystetstvo = Subject{
	Name:      "Музичне мистецтво",
	ShortName: "Муз. мистецтво",
}

var AnhliiskaMova = Subject{
	Name:      "Англійська мова",
	ShortName: "Англ. мова",
}

var ObrazotvorcheMystetstvo = Subject{
	Name:      "Образотворче мистецтво",
	ShortName: "Обр. мистецтво",
}

var FizychneVykhovannia = Subject{
	Name:      "Фізичне виховання",
	ShortName: "Фіз. виховання",
}

var YaDoslidzhuiuSvit = Subject{
	Name:      "Я досліджую світ",
	ShortName: "ЯДС",
}

var DyzainITekhnolohii = Subject{
	Name:      "Я досліджую світ",
	ShortName: "ЯДС",
}

var IndyvidualneZaniattia = Subject{
	Name:      "Дизайн і технології",
	ShortName: "Диз. і технології",
}

var Monday = Day{
	Name: "Понеділок",
	Subjects: [][]Subject{
		{LiteraturneChytannia},
		{Matematyka},
		{UkrainskaMova},
		{Informatyka},
		{MuzychneMystetstvo},
	},
}

var Tuesday = Day{
	Name: "Вівторок",
	Subjects: [][]Subject{
		{LiteraturneChytannia},
		{Matematyka},
		{AnhliiskaMova},
		{ObrazotvorcheMystetstvo},
		{FizychneVykhovannia},
	},
}

var Wednesday = Day{
	Name: "Середа",
	Subjects: [][]Subject{
		{YaDoslidzhuiuSvit},
		{Matematyka},
		{UkrainskaMova},
		{LiteraturneChytannia},
		{FizychneVykhovannia},
	},
}

var Thursday = Day{
	Name: "Четвер",
	Subjects: [][]Subject{
		{YaDoslidzhuiuSvit},
		{AnhliiskaMova},
		{Matematyka},
		{UkrainskaMova},
		{DyzainITekhnolohii},
		{IndyvidualneZaniattia},
	},
}

var Friday = Day{
	Name: "П'ятниця",
	Subjects: [][]Subject{
		{YaDoslidzhuiuSvit},
		{Matematyka},
		{UkrainskaMova, LiteraturneChytannia},
		{AnhliiskaMova},
		{FizychneVykhovannia},
	},
}

var Timetable = map[int]Day{
	1: Monday,
	2: Tuesday,
	3: Wednesday,
	4: Thursday,
	5: Friday,
}

var Bells = []SchoolBell{
	{
		Position: 1,
		Start:    "15:30",
		End:      "15:35",
	},
	{
		Position: 2,
		Start:    "15:40",
		End:      "15:45",
	},
	{
		Position: 3,
		Start:    "15:50",
		End:      "15:55",
	},
	{
		Position: 4,
		Start:    "16:00",
		End:      "16:05",
	},
	{
		Position: 5,
		Start:    "16:10",
		End:      "16:15",
	},
	{
		Position: 6,
		Start:    "16:20",
		End:      "16:25",
	},
}

//
//var Bells = []SchoolBell{
//	{
//		Position: 1,
//		Start:    "16:00",
//		End:      "16:40",
//	},
//	{
//		Position: 2,
//		Start:    "16:55",
//		End:      "17:35",
//	},
//	{
//		Position: 3,
//		Start:    "17:50",
//		End:      "18:30",
//	},
//	{
//		Position: 4,
//		Start:    "18:55",
//		End:      "19:35",
//	},
//	{
//		Position: 5,
//		Start:    "20:00",
//		End:      "20:40",
//	},
//	{
//		Position: 6,
//		Start:    "20:55",
//		End:      "21:35",
//	},
//}

//
//var Bells = []SchoolBell{
//	{
//		Position: 1,
//		Start:    "8:00",
//		End:      "8:40",
//	},
//	{
//		Position: 2,
//		Start:    "8:55",
//		End:      "9:35",
//	},
//	{
//		Position: 3,
//		Start:    "9:50",
//		End:      "10:30",
//	},
//	{
//		Position: 4,
//		Start:    "10:55",
//		End:      "11:35",
//	},
//	{
//		Position: 5,
//		Start:    "12:00",
//		End:      "12:40",
//	},
//	{
//		Position: 6,
//		Start:    "12:55",
//		End:      "13:35",
//	},
//}

var DayNames = map[int]string{
	0: "неділя",
	1: "понеділок",
	2: "вівторок",
	3: "середа",
	4: "четвер",
	5: "п'ятниця",
	6: "субота",
}
