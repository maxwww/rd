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

var Lesson = Subject{
	Name:      "Урок",
	ShortName: "Урок",
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
	Name:      "Дизайн і технології",
	ShortName: "Диз. і технології",
}

var IndyvidualneZaniattia = Subject{
	Name:      "Індивідуальне заняття",
	ShortName: "Індивідуальне заняття",
}

var Etyka = Subject{
	Name:      "Етика",
	ShortName: "Етика",
}

var PiznaiemoPryrodu = Subject{
	Name:      "Пізнаємо природу",
	ShortName: "Пізнаємо природу",
}

var Tekhnolohii = Subject{
	Name:      "Технології",
	ShortName: "Технології",
}

var UkrainskaLiteratura = Subject{
	Name:      "Українська література",
	ShortName: "Укр. література",
}

var Mystetstvo = Subject{
	Name:      "Ммистецтво",
	ShortName: "Мистецтво",
}

var Istoriia = Subject{
	Name:      "Історія",
	ShortName: "Історія",
}

var Zdorovia = Subject{
	Name:      "Здоров'я",
	ShortName: "Здоров'я",
}

var ZarubizhnaLiteratura = Subject{
	Name:      "Зарубіжна література",
	ShortName: "Зарубіжна література",
}

var Monday = Day{
	Name: "Понеділок",
	Subjects: [][]Subject{
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
	},
}

var Tuesday = Day{
	Name: "Вівторок",
	Subjects: [][]Subject{
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
	},
}

var Wednesday = Day{
	Name: "Середа",
	Subjects: [][]Subject{
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
	},
}

var Thursday = Day{
	Name: "Четвер",
	Subjects: [][]Subject{
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
	},
}

var Friday = Day{
	Name: "П'ятниця",
	Subjects: [][]Subject{
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
		{Lesson},
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
		Start:    "8:30",
		End:      "9:15",
	},
	{
		Position: 2,
		Start:    "9:25",
		End:      "10:10",
	},
	{
		Position: 3,
		Start:    "10:25",
		End:      "11:10",
	},
	{
		Position: 4,
		Start:    "11:20",
		End:      "12:05",
	},
	{
		Position: 5,
		Start:    "12:25",
		End:      "13:10",
	},
	{
		Position: 6,
		Start:    "13:20",
		End:      "14:05",
	},
	{
		Position: 7,
		Start:    "14:15",
		End:      "15:00",
	},
}

var DayNames = map[int]string{
	0: "неділя",
	1: "понеділок",
	2: "вівторок",
	3: "середа",
	4: "четвер",
	5: "п'ятниця",
	6: "субота",
}
