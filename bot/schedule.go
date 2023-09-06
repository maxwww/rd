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
		{FizychneVykhovannia},
		{Etyka},
		{Matematyka},
		{UkrainskaMova},
		{AnhliiskaMova},
		{PiznaiemoPryrodu},
	},
}

var Tuesday = Day{
	Name: "Вівторок",
	Subjects: [][]Subject{
		{Matematyka},
		{Tekhnolohii},
		{UkrainskaMova},
		{Informatyka},
		{AnhliiskaMova},
		{UkrainskaLiteratura},
	},
}

var Wednesday = Day{
	Name: "Середа",
	Subjects: [][]Subject{
		{FizychneVykhovannia},
		{Tekhnolohii},
		{AnhliiskaMova},
		{Matematyka},
		{UkrainskaMova},
	},
}

var Thursday = Day{
	Name: "Четвер",
	Subjects: [][]Subject{
		{Mystetstvo},
		{Informatyka},
		{Matematyka},
		{UkrainskaMova},
		{Istoriia},
		{Zdorovia},
	},
}

var Friday = Day{
	Name: "П'ятниця",
	Subjects: [][]Subject{
		{AnhliiskaMova},
		{FizychneVykhovannia},
		{ZarubizhnaLiteratura},
		{Matematyka},
		{UkrainskaLiteratura},
		{PiznaiemoPryrodu},
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
		Start:    "8:00",
		End:      "8:45",
	},
	{
		Position: 2,
		Start:    "8:55",
		End:      "9:40",
	},
	{
		Position: 3,
		Start:    "10:10",
		End:      "10:55",
	},
	{
		Position: 4,
		Start:    "11:05",
		End:      "11:50",
	},
	{
		Position: 5,
		Start:    "12:00",
		End:      "12:45",
	},
	{
		Position: 6,
		Start:    "12:55",
		End:      "13:40",
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
