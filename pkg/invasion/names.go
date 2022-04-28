package invasion

import (
	"fmt"
	"math/rand"
)

var (
	names = []string{"Absolutno", "Alasia", "Amadioha", "Amansinaya", "Anadolu", "Aniara", "Arcalís", "Atakoraka", "Axólotl",
		"Ayeyarwady", "Baekdu", "Bélénos", "Belel", "Berehynia", "Bibhā", "Bosona", "Bubup", "Buna", "Ceibo", "Chaophraya",
		"Chasoň", "Chechia", "Citadelle", "Citalá", "Cocibolca", "Dingolay", "Dìwö", "Diya", "Dofida", "Dombay", "Ebla", "Emiw",
		"Felixvarela", "Flegetonte", "Formosa", "Franz", "Funi", "Gakyid", "Gloas", "Gumala", "Hoggar", "Horna", "Hunahpú", "Hunor",
		"Illyrian", "Inquill", "Intan", "Irena", "Itonda", "Kalausi", "Kamuy", "Karaka", "Kaveh", "Koeia", "Koit", "Lerna", "Liesma", "Lionrock", "Lucilinburhuc", "Lusitânia", "Macondo", "Mago", "Mahsati", "Malmok", "Márohu", "Mazaalai", "Moldoveanu", "Mönch", "Montuno", "Morava", "Moriah", "Mouhoun", "Mpingo", "Muspelheim", "Naledi", "Násti", "Natasha", "Nenque", "Nervia", "Nikawiy", "Nosaxa", "Nushagak", "Nyamien", "Parumleo", "Petra", "Phoenicia", "Pincoya", "Pipoltr", "Poerava", "Rapeto", "Rosaliadecastro", "Sagarmatha", "Sāmaya", "Sansuna", "Shama", "Sharjah", "Sika", "Solaris", "Sterrennacht", "Stribor", "Taika", "Tangra", "Tapecue", "Tevel", "Timir", "Tislit", "Tojil", "Tuiren", "Tupã", "Tupi", "Uklun", "Uruk", "Xihe", "Gudja", "Guniibuu", "Imai", "La", "Paikauhale", "Toliman", "Alpherg", "Alruba", "Ashlesha", "Azmidi", "Bunda", "Elgafar", "Elkurud", "Fawaris", "Felis", "Fumalsamakah", "Heze", "Kraz", "Nahn", "Okab", "Piautos", "Tarf", "Ukdah", "Acamar", "Achernar", "Achird",
		"Acrab", "Acrux", "Acubens", "Adhafera", "Adhara", "Adhil", "Ain", "Ainalrami", "Aladfar", "Albaldah", "Albali", "Albireo", "Alchiba", "Alcor", "Alcyone", "Aldebaran", "Alderamin", "Aldhanab", "Aldhibah", "Aldulfin", "Alfirk", "Algedi", "Algenib", "Algieba", "Algol", "Algorab", "Alhena", "Alioth", "Aljanah", "Alkaid", "Alkalurops", "Alkaphrah", "Alkarab", "Alkes", "Almaaz", "Almach", "Alnair", "Alnasl", "Alnilam", "Alnitak", "Alniyat", "Alphard", "Alphecca", "Alpheratz", "Alrakis", "Alrescha", "Alsafi", "Alsciaukat", "Alsephina", "Alshain", "Alshat", "Altair", "Altais", "Alterf", "Aludra", "Alula",
		"Alula", "Alya", "Alzirr", "Ancha", "Angetenar", "Ankaa", "Anser", "Antares", "Arcturus", "Arkab", "Arkab", "Arneb", "Ascella", "Asellus", "Asellus", "Aspidiske", "Asterope", "Athebyne", "Atik", "Atlas", "Atria", "Avior", "Azelfafage", "Azha", "Barnard", "Baten", "Beemim", "Beid", "Bellatrix", "Betelgeuse", "Bharani", "Biham", "Botein", "Brachium", "Canopus", "Capella", "Caph", "Castor", "Castula", "Cebalrai", "Celaeno", "Cervantes", "Chalawan", "Chamukuy", "Chara", "Chertan", "Copernicus", "Cor", "Cujam", "Cursa", "Dabih", "Dalim", "Deneb", "Deneb",
		"Denebola", "Diadem", "Diphda", "Dschubba", "Dubhe", "Dziban", "Edasich", "Electra", "Elnath", "Eltanin", "Enif", "Errai", "Fafnir", "Fang", "Fomalhaut", "Fulu", "Furud", "Fuyue", "Gacrux", "Giausar", "Gienah", "Gomeisa", "Grumium", "Hadar", "Haedus", "Hamal", "Hassaleh", "Hatysa", "Helvetios", "Homam", "Iklil", "Intercrus", "Izar", "Jabbah", "Jishui", "Kaffaljidhma", "Kang", "Kaus", "Kaus", "Kaus", "Keid", "Khambalia", "Kitalpha", "Kochab", "Kornephoros", "Kurhah", "Lesath", "Libertas", "Lich", "Lilii", "Maasym", "Mahasim", "Maia", "Marfik", "Markab", "Markeb", "Marsic", "Matar", "Mebsuta", "Megrez", "Meissa", "Mekbuda", "Meleph", "Menkalinan", "Menkar", "Menkent", "Menkib", "Merak", "Merga", "Meridiana", "Merope", "Mesarthim", "Miaplacidus", "Mimosa", "Minchir", "Minelauva", "Mintaka", "Mira", "Mirach", "Miram", "Mirfak", "Mirzam", "Misam", "Mizar", "Mothallah", "Muliphein", "Muphrid", "Muscida", "Musica", "Naos", "Nashira", "Nekkar", "Nembus", "Nihal", "Nunki", "Nusakan", "Ogma", "Peacock", "Phact", "Phecda", "Pherkad", "Pipirima", "Pleione", "Polaris", "Polaris", "Polis", "Pollux", "Porrima", "Praecipua", "Prima", "Procyon", "Propus", "Proxima",
		"Ran", "Rasalas", "Rasalgethi", "Rasalhague", "Rastaban", "Regulus", "Revati", "Rigel", "Rigil", "Rotanev", "Ruchbah", "Rukbat", "Sabik", "Saclateni", "Sadachbia", "Sadalbari", "Sadalmelik", "Sadalsuud", "Sadr", "Saiph", "Salm", "Sargas", "Sarin", "Sceptrum", "Scheat", "Schedar", "Secunda", "Segin", "Seginus", "Sham", "Shaula", "Sheliak", "Sheratan", "Sirius", "Situla", "Skat", "Spica", "Sualocin", "Subra", "Suhail", "Sulafat", "Syrma", "Tabit", "Taiyangshou", "Taiyi", "Talitha", "Tania", "Tania", "Tarazed", "Taygeta", "Tegmine", "Tejat", "Terebellum", "Theemin", "Thuban", "Tiaki", "Tianguan", "Tianyi", "Titawin", "Tonatiuh", "Torcular", "Tureis", "Unukalhai", "Unurgunite", "Vega", "Veritate", "Vindemiatrix", "Wasat", "Wazn", "Wezen", "Xamidimura", "Xuange", "Yed", "Yed", "Yildun", "Zaniah", "Zaurak", "Zavijava", "Zhang", "Zibal", "Zosma", "Zubenelgenubi", "Zubenelhakrabi", "Zubeneschamali", "Larawag", "Ginan", "Wurren"}
)

func BuildAlienNamesArray(n int) []string {
	usedNames := make(map[string]struct{})
	resultNames := make([]string, n)
	for i := 0; i < n; i++ {
		rootNameInd := rand.Int31n(int32(len(names)))
		rootName := names[rootNameInd]
		if _, ok := usedNames[rootName]; !ok {
			usedNames[rootName] = struct{}{}
			resultNames[i] = rootName
			continue
		}
		for k := 2; ; k++ {
			name := fmt.Sprintf("%s-%d", rootName, k)
			if _, ok := usedNames[name]; !ok {
				usedNames[name] = struct{}{}
				resultNames[i] = name
				break
			}
		}
	}
	return resultNames
}
