package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/solorad/alien-invasion/pkg/invasion"
)

const (
	lineSize = 6
	// parameter for building a route in this direction. That will make some routes between cities one-directional
	addRoutePercentage = 90
)

var cityDictionary = []string{
	"Tokyo", "Delhi", "Shanghai", "Dhaka", "Cairo", "Beijing", "Mumbai", "Osaka", "Chongqing", "Karachi", "Istanbul", "Kinshasa", "Lagos", "Kolkata", "Manila", "Tianjin", "Guangzhou", "Lahore", "Bangalore", "Shenzhen", "Moscow", "Chennai", "Bogota", "Paris", "Jakarta", "Lima", "Bangkok", "Hyderabad", "Seoul", "Nagoya", "London", "Chengdu", "Nanjing", "Tehran", "Luanda", "Wuhan", "Ahmedabad", "Hangzhou", "Surat", "Suzhou", "Riyadh", "Shenyang", "Baghdad", "Dongguan", "Foshan", "Pune", "Santiago", "Madrid", "Haerbin", "Toronto", "Khartoum", "Johannesburg", "Singapore", "Dalian", "Qingdao", "Zhengzhou", "Barcelona", "Abidjan", "Yangon", "Fukuoka", "Alexandria", "Guadalajara", "Ankara", "Chittagong", "Melbourne", "Nairobi", "Hanoi", "Sydney", "Monterrey", "Changsha", "Brasilia", "Jiddah", "Urumqi", "Kunming", "Changchun", "Hefei", "Shantou", "Xinbei", "Kabul", "Ningbo", "Yaounde", "Rome", "Shijiazhuang", "Montreal", "Recife", "Kano", "Fortaleza", "Jaipur", "Nanning", "Medellin", "Ekurhuleni", "Douala", "Kozhikode", "Salvador", "Changzhou", "Xiamen", "Lucknow", "Casablanca", "Wenzhou", "Nanchang", "Malappuram", "Curitiba", "Ibadan", "Antananarivo", "Abuja", "Kampala", "Kumasi", "Faisalabad", "Bekasi", "Berlin", "Guiyang", "Busan", "Asuncion", "Campinas", "Wuxi", "Thrissur", "Dakar", "Mashhad", "Kochi", "Puebla", "Lanzhou", "Indore", "Durban", "Kanpur", "Sanaa", "Athens", "Milan", "Pyongyang", "Guayaquil", "Izmir", "Ouagadougou", "Lusaka", "Kiev", "Surabaya", "Nagpur", "Lisbon", "Zhongshan", "Dubai", "Caracas", "Depok", "Shizuoka", "Coimbatore", "Handan", "Port-au-Prince", "Huaian", "Algiers", "Cali", "Weifang", "Incheon", "Bamako", "Goiania", "Thiruvananthapuram", "Manchester", "Mbuji-Mayi", "Taibei", "Pretoria", "Zibo", "Shaoxing", "Lubumbashi", "Yantai", "Huizhou", "Chicago", "Sapporo", "Birmingham", "Bandung", "Vancouver", "Accra", "Tashkent", "Brazzaville", "Luoyang", "Patna", "Bhopal", "Damascus", "Mogadishu", "Brisbane", "Tangerang", "Tunis", "Beirut", "Nantong", "Medan", "Baku", "Belem", "Gujranwala", "Peshawar", "Manaus", "Sendai", "Maracaibo", "Rawalpindi", "Houston", "Barranquilla", "Agra", "Hohhot", "Taoyuan", "Baotou", "Kannur", "Liuzhou", "Visakhapatnam", "Vadodara", "Xuzhou", "Tijuana", "Esfahan", "Amman", "Daegu", "Naples", "Nashik", "Vijayawada", "Havana", "Mecca", "Brussels", "Multan", "Aleppo", "Putian", "Perth", "Yangzhou", "Hiroshima", "Baoding", "Bursa", "Minsk", "Conakry", "Kollam", "Rajkot", "Haikou", "Vienna", "Valencia", "Almaty", "Daqing", "Semarang", "Lianyungang", "Rabat", "Quito", "Hyderabad", "Lome", "Ludhiana", "Zhuhai", "Datong", "Quanzhou", "Adana", "Madurai", "Turin", "Matola", "Warsaw",
	"Hamburg", "Sharjah", "Bucharest", "Palembang", "Budapest", "Gaziantep", "Montevideo", "Meerut", "Raipur", "Phoenix", "Lyon", "Jiangmen", "Mosul", "Cixi", "Varanasi", "Xiangyang", "Shiraz", "Glasgow", "Novosibirsk", "Batam", "Stockholm", "Yinchuan", "Anshan", "Jamshedpur", "Yichang", "Srinagar", "Auckland", "Jilin", "Ulaanbaatar", "Tabriz", "Makassar", "Aurangabad", "Qinhuangdao", "Xining", "Muscat", "Monrovia", "Marseille", "Tiruppur", "Hengyang", "Calgary", "Qiqihaer", "Cordoba", "Suqian", "Kananga", "Karaj", "Anyang", "Philadelphia", "Rosario", "Daejon", "Munich", "Harare", "Onitsha", "Jodhpur", "Gaoxiong", "Medina", "N-Djamena", "Tegucigalpa", "Gwangju", "Yekaterinburg", "Kathmandu", "Edmonton", "Natal", "Ranchi", "Zhangjiakou", "Mandalay", "Jabalpur", "Huainan", "Asansol", "Kota", "Chaozhou", "Gwalior", "Allahabad", "Yiwu", "Nouakchott", "Amritsar", "Kharkiv", "Ottawa", "Zurich", "Basra", "Ganzhou", "Belgrade", "Homs", "Weihai", "Queretaro", "Mombasa", "Niamey", "Konya", "Jiaxing", "Copenhagen", "Cochabamba", "Dhanbad", "Kisangani", "Bucaramanga", "Kaifeng", "Adelaide", "Taizhong", "Rizhao", "Maceio", "Suweon", "Dallas", "Dongying", "Zunyi", "Zhanjiang", "Nanchong", "Joinville", "Qom", "Helsinki", "Liuan", "Porto", "Antalya", "Shiyan", "Prague", "Bareilly", "Liuyang", "Ad-Dammam", "Pointe-Noire", "Yingkou", "Sofia", "Kazan", "Tengzhou", "Aligarh", "Ahvaz", "Florianopolis", "Tanger", "Freetown", "Maoming", "Fes", "Moradabad", "Suzhou", "Uyo", "Mysore", "Dublin", "Astana", "Ruian", "Mwanza", "Durg-Bhilainagar", "Barquisimeto", "Jieyang", "Chelyabinsk", "Zhuzhou", "Baoji", "Maracay", "Bhubaneswar",
	"Chifeng", "Puning", "Lilongwe", "Jinhua", "Mendoza", "Kigali", "Bogor", "Huaibei", "Merida", "Tiruchirappalli", "Islamabad", "Benxi", "Jinzhou", "Chandigarh", "Bukavu", "Abomey-Calavi", "Liupanshui", "Omsk", "Nnewi", "Tripoli", "Guilin", "Amsterdam", "Tasikmalaya", "Haifa", "Binzhou", "Pizhou", "Quetta", "Mexicali", "Krasnoyarsk", "Hubli-Dharwad", "Kaduna", "Samara", "Guwahati", "Aba", "Luohe", "Salem", "Aguascalientes", "Ufa", "Bujumbura", "Maputo", "Rostov-on-Don", "Cologne", "Yueqing", "Saharanpur", "Shimkent", "Yongin", "Xinxiang", "Bazhong", "Xiongan", "Zaozhuang", "Cuernavaca", "Wenling", "Panjin", "Chihuahua", "Jalandhar", "Goyang", "Fuyang", "Yerevan", "Siliguri", "Managua", "Ma'anshan", "Bishkek", "Shangrao", "Tbilisi", "Hargeysa", "Cartagena", "Diyarbakir", "Perm", "Zhaoqing", "Voronezh", "Bobo-Dioulasso", "Lille", "Samarinda", "Oslo", "Kermanshah", "Leshan", "Solapur", "Changwon", "Mersin", "Antwerp", "Kirkuk", "Dezhou", "Toulouse", "Chenzhou", "Aden", "Aracaju", "Huzhou", "Ikorodu", "Teresina", "Asmara", "Yueyang", "Xuchang", "Marrakech", "Tampico", "Denpasar", "Austin", "Tshikapa", "Changshu", "Saltillo", "Padang", "Rotterdam", "Nyala", "Qujing", "Odesa", "Warangal", "Ilorin", "Valparaiso",
	"Kayseri", "Cancun", "Volgograd", "Xintai", "Blantyre-Limbe", "Songkhla", "Bordeaux", "Krasnodar", "Nonthaburi", "Guiping", "Morelia", "Dehradun", "Hamah", "Dushanbe", "Laiwu", "Agadir", "Jiujiang", "Zhucheng", "Mudanjiang", "Chengde", "Jerusalem", "Misratah", "Cucuta", "Khulna", "Arequipa", "Dnipro", "Nice", "Owerri", "Southampton", "Guigang", "Warri", "Veracruz", "Jos", "Rajshahi", "Seongnam", "Leiyang", "Jacksonville", "Yangjiang", "Reynosa", "Bangui", "Najaf", "Bengbu", "Zhangzhou", "Villahermosa", "Sylhet", "Nampula", "Tianmen", "Charlotte", "Antipolo", "Hengshui", "Xinyu", "Linfen", "Oran", "Columbus", "Lubango", "Deyang", "Ulsan", "Liverpool", "Londrina", "Taiz", "Bergamo", "Malang", "Jiangyin", "Dandong", "Concepcion", "Hermosillo", "Liling", "Guntur", "Baishan", "Bahawalpur", "Donetsk", "Indianapolis", "Trujillo", "Changzhi", "Bhiwandi", "Tyumen", "Liaoyang", "Ashgabat", "Soshanguve", "Erbil", "Puducherry", "Changde", "Shangqiu", "Culiacan", "Firozabad", "Xingtai", "Bogra", "Tainan", "Fuxin", "Cabinda", "Huangshi", "Umuahia", "Hufuf-Mubarraz", "Libreville", "Palermo", "Yibin", "Banghazi", "Luzhou", "Ipoh", "Kuerle", "Xinghua", "Saratov", "Cherthala", "Valencia", "Quzhou", "Xinyang", "Yongzhou", "Winnipeg", "Yangquan", "Xiaogan", "Bucheon", "Orumiyeh", "Maiduguri", "Enugu", "Eskisehir", "Zhuji", "Heze", "Huaihua", "Tianshui", "Thessaloniki", "Bologna", "Haicheng", "Huludao", "Sorocaba", "Bozhou", "Bikaner",
	"Nottingham", "Niigata", "Xalapa", "Kunshan", "Liaocheng", "Nellore", "Gebze", "Jincheng", "Frankfurt", "Taixing", "Lokoja", "Seattle", "Langfang", "Vereeniging", "Jiaozuo", "Dasmarinas", "Kottayam", "Ar-Rayyan", "Sulaimaniya", "Amravati", "Zhumadian", "Benguela", "Gorakhpur", "Hamilton", "Shaoguan", "Meishan", "Gaomi", "Hanchuan", "Krakow", "Bunia", "Danyang", "Muzaffarnagar", "Zanzibar", "Ansan", "Cuttack", "Anqing", "Denver", "Anqiu", "Banjarmasin", "Kayamkulam", "Zigong", "Gaza", "Qingyuan", "Linhai", "Erduosi-Ordoss", "Belgaum", "Oshogbo", "Zaria", "Malegaon", "Shaoyang", "Malanje", "Zarqa", "Sheffield", "Kumamoto", "Yan'an", "Maturin", "Tongliao", "Jiamusi", "Yanji", "Zaragoza", "Sialkot", "Mangalore", "Kitwe", "Tongling", "Wuzhou", "Dengzhou", "Ibb", "Rasht", "Yuncheng", "Tirupati", "Zaporizhzhya", "Merca", "Dazhou", "Panzhihua", "Sargodha", "Cheongju", "Bhavnagar", "Okayama", "Yuxi", "Bacoor", "Lviv", "Dongtai", "Washington", "Akure", "Jammu", "Longyan", "Yuyao", "Al-Hudaydah", "Bali", "Florence", "Zhangjiagang", "Cotonou", "Salta", "Goma", "Huambo", "Vientiane", "Kurnool", "Amara", "Taicang", "Uberlandia", "Gulbarga", "Taif", "Seville", "Tamale", "Bristol", "Buraydah", "Celaya", "Durgapur", "Jundiai", "Boston", "Cheonan", "Ankang", "Tolyatti", "Pontianak", "Nantes", "Zhoukou",
	"Padova", "Manama", "Loudi", "Renqiu", "Bijie", "Lattakia", "Ezhou", "Sokoto", "Zagreb", "Liege", "Nashville", "Hezhou", "Tongxiang", "Santiago", "Qinzhou", "Samsun", "Longhai", "Tabuk", "Denizli", "Genoa", "Puyang", "Zhoushan", "Miluo", "Macao", "Portland", "Lodz", "Jhansi", "Herat", "Haimen", "Jeonju", "Sanya", "Neijiang", "Erode", "Detroit", "As-Suways", "Wuhai", "Jamnagar", "Uvira", "Jubayl", "Durango", "Chiclayo", "Nasiriyah", "Hegang", "Izhevsk", "Doha", "Memphis", "Enshi", "Chongjin", "Jambi", "Sihui", "Bauchi", "Irkutsk", "Bulawayo", "Safaqis", "Bissau", "Wroclaw", "Raurkela", "Barnaul", "Venezia", "Khabarovsk", "Belfast", "Duesseldorf", "Anshun", "Bacolod", "Verona", "Kolhapur", "Stuttgart", "Jinzhong", "Al-Ain", "Patiala", "Bilaspur", "Abakaliki", "Sanhe", "Kuching", "Hanzhong", "Calabar", "Ajmer", "Comilla", "Ujjain", "Seregno", "Hillah", "Geneva", "Colombo", "Xinmi", "Riga", "Gothenburg", "Ulyanovsk", "Zahedan", "Bari", "Mbeya", "Jingjiang", "Mathura", "Cenxi", "Brighton", "Tlaxcala", "Siping", "Louisville", "Agartala", "Yaroslavl", "Shishi", "Imphal", "Fuqing", "Vladivostok", "Daye", "Skopje", "Likasi", "Tomsk", "Shuozhou", "Cuito", "Samarkand", "Sanliurfa", "Ogbomosho", "Beira", "Sangli", "Makhachkala", "Xiantao", "Leipzig", "Dengfeng", "Cabimas", "Qitaihe", "Sanmenxia", "Tonghua", "Ibague", "Thoothukkudi", "Cuiaba", "Xinyi", "Kingston", "Zhaodong", "Malaga", "Udaipur", "Hamadan", "Djibouti", "Dongyang",
	"Bouake", "Dortmund", "Zaoyang", "Vellore", "Jalgaon", "Ardabil", "Bloemfontein", "Hebi", "Gaya", "Milwaukee", "Catania", "Kagoshima", "Dresden", "Yazd", "Jiaozhou", "Gaozhou", "Toulon", "Tirunelveli", "Uige", "Oujda", "Pereira", "Essen", "Matamoros", "Kitchener", "Imus", "Cangzhou", "Karbala", "Weinan", "Al-Mansurah", "Guangyuan", "Bamenda", "Panipat", "Kahramanmaras", "Ndola", "Feicheng", "Tepic", "Orenburg", "Meknes", "Kemerovo", "Arak", "Bremen", "Gomel", "Baltimore", "Tongchuan", "Sakarya", "Basel", "Al-Mukalla", "Mekele", "Larkana", "Albuquerque", "Hosur", "Anyang", "Villavicencio", "Rupganj", "Leicester", "Shouguang", "Irbid", "Valledupar", "Rajahmundry", "Yongkang", "Utrecht", "Tucson", "Sukkur", "Haining", "Abeokuta", "Yanshi", "Chaoyang", "Rustenburg", "Taishan", "Calamba", "Ambon", "Diwaniyah", "Rayong", "Chuzhou", "Gombe", "Qingzhou", "Changning", "Longkou", "Loum", "Novokuznetsk", "Mesa", "Edinburgh", "Pingdu", "Dali", "Kerman", "Kikwit", "Bellary", "Mbouda", "Al-Raqqa", "Berbera", "Gaoyou", "Sheikhupura", "Namangan", "Hamhung", "Gongyi", "Balashikha", "Zinder", "Van", "Heyuan", "Vilnius", "Laixi", "Ryazan", "Fresno", "Shizuishan", "Pietermaritzburg", "Chuxiong", "Changge", "Xingning", "Hannover", "Shangyu", "Surakarta", "Kabinda", "Astrakhan", "Grenoble", "Atlanta", "Kakinada", "Kalasin", "Sacramento", "Davanagere", "Qianjiang", "Muzaffarpur", "Ziyang", "Mataram", "Ar-Rusayfah", "Kandahar", "Changyi",
	"Songyuan", "Poznan", "Kuantan", "Eslamshahr", "Duisburg", "Guang-an", "Penza", "Jinjiang", "Utsunomiya", "Begusarai", "Arusha", "Kirov", "Barisal", "Ado-Ekiti", "Khon-Kaen", "Purnia", "Tanta", "Jianyang", "Bournemouth", "Kismaayo", "Shuangyashan", "Nurenberg", "Malatya", "Cheboksary", "Tirana", "Kolwezi", "Nanping", "Jingmen", "Murcia", "Matsuyama", "Yongcheng", "Karaganda", "Zhongxiang", "Katsina", "Lipetsk", "Macapa", "Chaohu", "Zhuanghe", "Latur", "Douai-Lens", "Xuancheng",
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("city num argument must be provided")
	}
	cityNumStr := os.Args[1]
	cityNum, err := strconv.Atoi(cityNumStr)
	if err != nil || cityNum < 1 {
		log.Fatalf("city num argument must be valid positive integer")
	}
	// 0. we use random function a lot, so put it here
	rand.Seed(time.Now().UnixNano())
	cities := invasion.GetRandomNames(cityNum, cityDictionary)
	cityNamesMesh := fillCityMesh(cities, cityNum)
	for i := 1; i < len(cityNamesMesh)-1; i++ {
		line := cityNamesMesh[i]
		for j := 1; j < len(line)-1; j++ {
			if cityNamesMesh[i][j] == "" {
				continue
			}
			cityLine := []string{cityNamesMesh[i][j]}
			cityLine = addNeighbour("north", cityNamesMesh[i-1][j], cityLine)
			cityLine = addNeighbour("east", cityNamesMesh[i][j+1], cityLine)
			cityLine = addNeighbour("south", cityNamesMesh[i+1][j], cityLine)
			cityLine = addNeighbour("west", cityNamesMesh[i][j-1], cityLine)
			fmt.Println(strings.Join(cityLine, " "))
		}
	}
}

func addNeighbour(direction, cityName string, cityLine []string) []string {
	val := rand.Int31n(100)
	if val > addRoutePercentage {
		// don't add route even though it might present by a mesh
		return cityLine
	}
	if cityName != "" {
		return append(cityLine, fmt.Sprintf("%s=%s", direction, cityName))
	}
	return cityLine
}

func fillCityMesh(cities []string, cityNum int) [][]string {
	cityNamesMesh := make([][]string, 0)
	// put an empty line in a mesh on borders for protection
	cityNamesMesh = append(cityNamesMesh, make([]string, lineSize))
	totalCities := 0
	for {
		line := make([]string, lineSize)
		cityNamesMesh = append(cityNamesMesh, line)
		for i := 1; i < lineSize-1; i++ {
			// we skip position in 50% and in 50% - put a city in it
			putCity := rand.Intn(2)
			if putCity < 1 {
				continue
			}
			line[i] = cities[totalCities]
			totalCities++
			if totalCities == cityNum {
				cityNamesMesh = append(cityNamesMesh, make([]string, lineSize))
				return cityNamesMesh
			}
		}
	}
}
