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

var FizychnaKultura = Subject{
	Name:      "Фізична культура",
	ShortName: "Фіз. культура",
}

var IstoriiaUkrainy = Subject{
	Name:      "Історія України",
	ShortName: "Іст. Укр.",
}

var UkrainskaMova = Subject{
	Name:      "Українська мова",
	ShortName: "Укр. мова",
}

var AnhliiskaMova = Subject{
	Name:      "Англійська мова",
	ShortName: "Англ. мова",
}

var UkrainskaLiteratura = Subject{
	Name:      "Українська література",
	ShortName: "Укр. література",
}

var Alhebra = Subject{
	Name:      "Алгебра",
	ShortName: "Алгебра",
}

var Heometriia = Subject{
	Name:      "Геометрія",
	ShortName: "Геометрія",
}

var Informatyka = Subject{
	Name:      "Інформатика",
	ShortName: "Інформатика",
}

var Biolohiia = Subject{
	Name:      "Біологія",
	ShortName: "Біологія",
}

var Mystetstvo = Subject{
	Name:      "Мистецтво",
	ShortName: "Мистецтво",
}

var Fizyka = Subject{
	Name:      "Фізика",
	ShortName: "Фізика",
}

var Khimiia = Subject{
	Name:      "Хімія",
	ShortName: "Хімія",
}

var HromadianskaOsvita = Subject{
	Name:      "Громадянська освіта",
	ShortName: "Гром. освіта",
}

var Heohrafiia = Subject{
	Name:      "Географія",
	ShortName: "Географія",
}

var TrudoveNavchannia = Subject{
	Name:      "Трудове навчання",
	ShortName: "Трудове",
}

var VsesvitniaIstoriia = Subject{
	Name:      "Всесвітня Історія",
	ShortName: "Всесв. Іст.",
}

var ZarubizhnaLiteratura = Subject{
	Name:      "Зарубіжна література",
	ShortName: "Зарубіжна література",
}

var ZdoroviaBezpekaTaDobrobut = Subject{
	Name:      "Здоров'я, безпека та добробут",
	ShortName: "ЗБД",
}

var Monday = Day{
	Name: "Понеділок",
	Subjects: [][]Subject{
		{FizychnaKultura},
		{IstoriiaUkrainy},
		{AnhliiskaMova},
		{UkrainskaMova},
		{UkrainskaLiteratura},
		{Alhebra},
		{Informatyka},
	},
}

var Tuesday = Day{
	Name: "Вівторок",
	Subjects: [][]Subject{
		{Biolohiia},
		{Mystetstvo},
		{Fizyka},
		{Heometriia},
		{FizychnaKultura},
		{Khimiia},
		{HromadianskaOsvita},
	},
}

var Wednesday = Day{
	Name: "Середа",
	Subjects: [][]Subject{
		{AnhliiskaMova},
		{UkrainskaMova},
		{Alhebra},
		{Heohrafiia},
		{UkrainskaLiteratura},
		{TrudoveNavchannia},
		{Khimiia},
	},
}

var Thursday = Day{
	Name: "Четвер",
	Subjects: [][]Subject{
		{Fizyka},
		{VsesvitniaIstoriia},
		{AnhliiskaMova},
		{ZarubizhnaLiteratura},
		{UkrainskaMova},
		{Heometriia},
		{FizychnaKultura},
	},
}

var Friday = Day{
	Name: "П'ятниця",
	Subjects: [][]Subject{
		{ZarubizhnaLiteratura},
		{AnhliiskaMova},
		{Alhebra},
		{UkrainskaMova},
		{Heohrafiia},
		{Biolohiia},
		{ZdoroviaBezpekaTaDobrobut},
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
