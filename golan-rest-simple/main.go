package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type carrera struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
	Campus int    `json:"Campus"`
}

type todas_las_tareas []carrera

var carreras = todas_las_tareas{
	{
		ID:     1,
		Nombre: "Ing. Sistemas computacionales",
		Campus: 1,
	},
	{
		ID:     2,
		Nombre: "Ing. Industrial",
		Campus: 1,
	},
	{
		ID:     3,
		Nombre: "Ingenieria en tecnologias de la informacion y comunicaciones",
		Campus: 2,
	},
	{
		ID:     4,
		Nombre: "Ing. en gestion empresarial",
		Campus: 1,
	},
	{
		ID:     5,
		Nombre: "Ing. en Electronica",
		Campus: 2,
	},
	{
		ID:     6,
		Nombre: "Electromecanica",
		Campus: 2,
	},
	{
		ID:     7,
		Nombre: "Mecatronica",
		Campus: 1,
	},
	{
		ID:     8,
		Nombre: "Ing. en logistica",
		Campus: 2,
	},
}

func obtenerCarreras(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.MarshalIndent(carreras, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Establecer el encabezado Content-Type en application/json
	w.Header().Set("Content-Type", "application/json")
	// Escribir la respuesta JSON en el cuerpo de la respuesta HTTP
	w.Write(jsonData)
}

func obtener_carrera(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["ID"])

	if err != nil {
		fmt.Fprint(w, "ID inválido")
		return
	}

	var materias []string
	for _, carrera := range carreras {
		if carrera.ID == taskID {
			materias = obtenerMateriasPorCarrera(carrera.Nombre)
			break
		}
	}

	if len(materias) == 0 {
		fmt.Fprint(w, "Carrera no encontrada")
		return
	}

	respuesta := struct {
		Carrera  string   `json:"Carrera"`
		Materias []string `json:"Materias"`
	}{
		Carrera:  obtenerNombreCarreraPorID(taskID),
		Materias: materias,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func obtenerMateriasPorCarrera(nombreCarrera string) []string {
	// Aquí puedes implementar la lógica para obtener las materias de una carrera específica
	// Por ahora, proporcionaré un ejemplo simple
	switch nombreCarrera {
	case "Ing. Sistemas computacionales":
		return []string{"Calculo diferencial", "Matematicas discretas", "Fundamentos de programacion", "Taller de administracion", "Fundamentos de investigacion",
			"Taller de etica", "Calculo integral", "Algebra lineal", "Programacion orientada a objetos", "Contabilidad financiera", "Qumica", "Probabilidad y estadistica",
			"Calculo vectorial", "Investigacion de operaciones", "Estructura de datos", "Cultura empresarial", "Desarrollo Sustentable", "Fisica general",
			"Ecuaciones diferenciales", "Fundamentos de base de datos", "Topicos Avanzados de programacion", "Simulacion", "Metodos numericos", "Principios electricos y aplic. digitales",
			"Fundamentos de ingenieria de software", "Taller de base de datos", "Sistemas operativos", "Graficacion", "Fundamentos de telecomunicaciones", "Arquitectura de computadoras",
			"Ingenieria de software", "Administracion de base de datos", "Lenguajes y automatas I", "Taller de sistemas operativos", "Redes de computadoras", "Lenguajez de interfaz",
			"Gestion de proyectos de software", "Lenguajes y automatas II", "Taller de investigacion I", "Conmutacion y enrutamiento de redes de datos", "Sistemas programables",
			" Taller de investigacion II", "Administracion de redes", "Programacion logica y funcional", "Programacion web", "Inteligencia arrtificial"}

	case "Ing. Industrial":
		return []string{"Fundamentos de investigacion", "Calculo diferencial", "Desarrollo humano", "Fundamentos de gestion empresarial", "Dinamica social", "Fundamentos de quimica",
			"Software de aplicacion ejecutivo", "Calculo integral", "Cont. orientada a los negoc.", "Legislacion laboral", "Taller de etica", "Fundamentos de fisica",
			"Marco legal de las organizaciones", "Probabilidad y estad. descrip.", "Costos empresariales", "Habilidades directivas I", "Economia empresarial", "Algebra lineal",
			"INSTR. de PRES. e,presarial", "Estadistica inferencial I", "Ingenieria de procesos", "Habilidades directivas II", "Entorno macroeconomico", "Invest. de operaciones",
			"Finanzas en las organizaciones", "Estadistica inferencial II", "Gestion de la produccion I", "Gestion de capital humano", "Desarrollo sustentable", "Mercadotecnia",
			"Ingenieria economica", "El emprendedor y la innovacion", "Gestion de la produccion II", "Admon. de la salud y seg. ocupacional", "Calidad aplicada a la gestion empresarial",
			"Sist. de inf. de la MKT", "Plan de negocios", "Taller de investigacion I", "Diseño organizacional", "Mercadotecnia electronica", "Taller de investigacion II",
			"Gestion estrategica", "Cadena de suministros"}

	case "Ingenieria en tecnologias de la informacion y comunicaciones":
		return []string{"Cálculo Diferencial", "Fundamentos de programación", "Matemáticas discretas", "Introducción a las TIC'S", "Taller de ética", "Fundamentos de investigación",
			"Cálculo integral", "Programación orientada a objetos", "Matematicas discretas II", "Probabilidad y estadística", "Contabilidad y costos", "Administración gerencial",
			"Matematicáticas aplicadas a comunicaciones", "Estructura y organización de datos", "Álgebra lineal", "Fundamentos de bases de datos", "Electricidad y magnetimo", "Desarrollo sustentable",
			"Análisis de señales y sistemas de Comunicación", "Programación II", "Matemáticas para la toma de decisiones", "Taller de base de datos", "Circuitos eléctricos y electrónicos", "Ingeniería de software",
			"Redes de computadoras", "Tecnologías inalámbricas", "Desarrollo de emprendedores", "Taller de investigación I", "Sistemas operativos I", "Programación web",
			" Redes emergentes", "Desarrollo de aplicaciones para dispositivos moviles", "Ingeniería del conocimiento", "Taller de investigación II", "Sistemas operativos II", "Negocios electrónicos I",
			"Administración y seguridad de redes", "Interación humano computadora", "Auditoría en tecnologias de información", "Negocios electronicos II", "Residencias"}
	case "Ing. en gestion empresarial":
		return []string{"Fundamentos de la investigación", "Cálculo diferencial", "Desarrollo humano", "Fundamentos de gestión empresarial", "Dinámica social", "Fundamentos de química",
			"Software de aplicación ejecutivo", "Cálculo integral", "Contabilidad orientada a los negocios", "Legislación laboral", "Taller de ética", "Fundamentos de física",
			"Marco legal de las organizaciones", "Probabilidad y estadística descriptiva", "Costos empresariales", "Habilidades directivas I", "Economía empresarial", "Álgebra lineal",
			"Instrumentos de Presupuestación empresarial", "Estadística inferencial I", "Ingeníeria de procesos", "Habilidades directivas II", "Entorno macroeconómico", "Investigación de operaciones",
			"Finanza en las organizaciones", "Estadística inferencial II", "Gestíon de la producción I", "Gestión del capital humano", "Desarrollo sustentable", "Mercadotecnia",
			"Ingeniería económica", "El emprendedor y la innovación", "Gestión de la producción II", "Administración de la salud y seguridad ocupacional", "Calidad aplicada a la gestión empresarial", "Sistemas de información de mercadotecnia",
			"Plan de negocios", "Taller de investigación I", "Diseño organizacional", "Mercadotecnia electrónica", "Taller de investigación II",
			" Gestión estrategica", "Cadena de suministros", "Residencia profesional"}
	case "Ing. en Electronica":
		return []string{"Cálculo diferencial", "Mecánica clasica", "Química", "Taller de ética", "Fundamentos de investigación", "Comunicación humana",
			"Cálculo integral", "Probabilidad y estadística", "Desarrollo sustentable", "Mediciones electricas", "Tópicos selectos de física", "Desarrollo humano",
			"Cálculo vectorial", "Electromagnetismo", "Álgebra lineal", "Física de semiconductores", "Programación estructurada",
			"Ecuaciones diferenciales", "Circuitos eléctricos I", "Marco legal de la empresa", "Análisis numérico", "Diseño digital", "Programación visual",
			"Circuitos eléctricos II", "Diodos y transitores", "Teoría electromagnética", "Máquinas eléctricas", "Diseño digital con VHDL", "Desarrollo profesional",
			"Control I", "Diseño con transitores", "Fundamentos financieros", "Microcontroladores", "Taller de investigación I", "Control II", "Amplificadores operacionales", "Instrumentción", "Optoeléctronica", "Introduccion a las telecomunicaciones",
			"Taller de investigación II", "Control digital", "Controladores lógicos programables", "Eléctronica de potencia", "Administración gerencial", "Desarrollo y evaluación de proyectos", "Residencia profesional"}

	case "Electromecanica":
		return []string{"Taller de ética, Cálculo diferencial", "Introducción a la programación", "Desarrollo sustentable", "Química", "Fundamentos de investigación",
			"Estática", "Cálculo integral", "Álgebra lineal", "Metrología y normalización", "Tecnología de los materiales", "Dibujo electromecánico",
			"Dinámica", "Cálculo vectorial", "Procesos de manufactura", "Electricidad y magnetismo", "Mecánica de materiales", "Probabilidad y etadística",
			"Análisis y sintasis de mecanismos", "Ecuaciones diferenciales", "Termodínamica", "Análisis de circuitos eléctricos de CD", "Mecánica de fluidos", "Electrónica analógica",
			"Diseño de elementos de máquinas", "Diseño e ingeniería asistidos por computadora", "Transferencia de calor", "Análisis de circuitos eléctricos de CA", "Sistemas y máquinas de fluidos", "Electrónica digital",
			"Máquinas y equipos térmicos I", "Ahorro de energía", "Instalaciones eléctricas", "Máquinas Eléctricas", "Administración y Técnicas de Mantenimiento", "Taller de Investigación I",
			"Máquinas y Equipos Térmicos II", "Sistemas Eléctricos de Potencia", "Controles Eléctricos", "Ingeniería de Control Clásico", "Taller de Investigación II", "Refrigeración y Aire Acondicionado",
			"Subestaciones Eléctricas", "Sistemas Hidráulicos y Neumáticos de Potencia", "Formulación y Evaluación de Proyectos"}

	case "Mecatronica":
		return []string{"QUIMICA", "CALCULO DIFERENCIAL", "TALLER DE ETICA", "DIBUJO ASISTIDO POS COMPUTADORA", "METROLOGIA Y NORMALIZACION", "FUNDAMENTOS DE INVESTIGACION",
			"CALCULO INTEGRAL", "ALGEBRA LINEL", "CIENCIAS E INGENIERIAS DE MATERIALES", "PROGRAMACION BASICA", "ESTADISTICA Y CONTROL DE CALIDAD", "ADMINISTRACION Y CONTABILIDAD",
			"CALCULO VECTORIAL", "PROCESOS DE FABRICACION", "ELECTROMAGNETISMO", "ESTATICA", "METODOS NUMERICOS", "DESARROLLO SUSTENTABLE",
			"ECUACIONES DIFERENCIALES", "FUNDAMENTOS DE TERMODINAMICA", "MECANICA DE MATERIALES", "DINAMICA", "ANALISIS DE CIRCUITOS ELECTRICOS",
			"MAQUINAS ELECTRICAS", "ELECTRONICA ANALOGICA", "MECANICANISMOS", "ANALISIS DE DLUIDOS", "TALLER DE INVESTIGACION 1", "ACTIVIDADES COMPLEMENTARIAS",
			"ELECTRONICA DE POTENCIA APLICADA", "INSTRUMENTACION", "DISEÑO DE ELEMENTOS MECANICOS", "ELECTRONICA DIIGITAL", "VIBRACIONES MECANICAS", "TALLER DE INVESTIGACION 2",
			"DINAMICA DE SISTEMAS", "MANUFACTURA AVANZADA", "CIRCUITOS HIDRAULICOS Y NEUMATICOS", "MANTENIMIENTO", "MICROCONTROLADORES", "PROGRAMMACION AVANZADA",
			"CONTROL", "FORMULACION Y EVALUACION DE PROYECTOS", "CONTROLADORES LOGICOS PROGRAMABLES", "SERVICIO SOCIAL",
			"ROBOTICA", "RESIDENCIA PROFESIONAL", "ESPECIALIDAD"}

	case "Logistica":
		return []string{"INTRODUCCION A LA INGENIERIA EN LOGISTICA", "CALCULO DIFERENCIAL", "QUIMICA", "FUNDAMENTOS DE ADMINISTRACION", "DIBUJO ASISTIDO POR COMPUTADORA", "ECONOMIA",
			"TALLER DE ETICA", "CALCULO INTEGRAL", "PROBABILIDAD Y ESTADISTICA", "DESARROLLO HUMANO Y ORGANIZACIONAL", "FUNDAMENTOS DE INVESTIGACION", "CONTABILIDAD Y COSTOS",
			"CADENA DE SUMINISTRO", "ALGEBRA LINEAL", "ESTADISTICA INFERENCIAL 1", "FUNDAMENTOS DE DERECHO", "MECANICA CLASICA", "FINANZAS",
			"COMPRAS", "TIPOLOGIA DEL PRODUCTO", "ESTADISTICA INFERENCIAÑ 2", "ENTORNO ECONOMICO", "TOPICOS DE INGENIERIA MECANICA", "BASES DE DATOS",
			"ALMACENES", "INVENTARIOS", "INVESTIGACION DE OPERACIONES 1", "HIGIENE Y SEGURIDAS", "PROCESOS DE FABRICACION Y MANEJO DE MATERIALES", "MERCADOTECNIA",
			"TRAFICO Y TRANSPORTE", "CULTURA DE CALIDAD", "INVESTIGACION DE OPERACIONES 2", "DESARROLLO SUSTENTABLE", "TALLER DE INVESTIGACION 1", "EMPAQUE,ENVASE Y EMBALAJE", "ACTIVIDADES EXTRAESCOLARES",
			"SERVICIO AL CLIENTE", "PROGRAMACION DE PROCESOS PRODUCTIVOS", "MODELOS DE SIMULACION Y LOGISTICA", "LEGISLACION ADUANERA", "TALLER DE INVESTIGACION 2", "INGENIERIA ECONOMICA",
			"INNOVACION", "COMERCIO INTERNACIONAL", "FORMULACION Y EVALUACION DE PROYECTOS", "GEOGRAFIA PARA EL TRANSPORTE", "SERVICIO SOCIAL",
			"RESIDENCIA PROFESIONAL", "ESPECIALIDAD", "GESTION DE PROYECTOS"}
	default:
		return nil
	}
}

func obtener_especialidad(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["ID"])

	if err != nil {
		fmt.Fprint(w, "ID inválido")
		return
	}

	especialidad, err := strconv.Atoi(vars["numero"])
	if err != nil {
		fmt.Fprint(w, "Especialidad inválida")
		return
	}

	// Obtener el nombre de la carrera por ID
	nombreCarrera := obtenerNombreCarreraPorID(taskID)
	if nombreCarrera == "" {
		fmt.Fprint(w, "Carrera no encontrada")
		return
	}

	// Definir las materias de especialidad para cada carrera
	var materias []string
	switch taskID {
	case 1: // Ing. Sistemas computacionales
		switch especialidad {
		case 1:
			materias = []string{"Topicos para el despliegue de aplicaciones", "Inteligencia de negocios", " Servicios en la nube", "Desarrollo de aplicaciones moviles",
				"Inteligencia de negocios II"}
		case 2:
			materias = []string{"Machine Learning", "Linguistica Computacional", "Softcomputing", "Procesamiento digital de señales", "Introduccion a la computacion cuantica"}
		default:
			fmt.Fprint(w, "No hay mas especialidades")
			return
		}
	case 2: // Ing. Industrial
		switch especialidad {
		case 1:
			materias = []string{"Direccion y potencializacion del talento humano", "Gestion del conocimiento y capital intelectual", "Gestion del cambio",
				"Estrategias financieras aplicadas a la gestion del capital humano", "Temas selectos de calidad"}
		case 2:
			materias = []string{"Sistemas neumaticos e hidraulicos", "Diseño asistido por computadora", "Temas selectos de manufactura", "Optimizacion de sistemas de manufactura",
				" Medicion y mejoramiento de la productividad", "Robotica industrial", "Manufactura flexible"}
		default:
			fmt.Fprint(w, "No hay mas especialidades")
		}

	case 3: // Ingenieria en tecnologias de la informacion y comunicaciones
		switch especialidad {
		case 1:
			materias = []string{"Seguridad informatica", "Administración de servidores", "Gestion de sistemas voIP",
				"Estrategias de gestión de servicios de TI", "Virtualización & loT"}
		}
	case 4: // Ing. en gestion empresarial
		switch especialidad {
		case 1:
			materias = []string{"Seminiario de innovación de procesos", "Calidad aplicada a la gestión empresarial II", "Seminario de consultoría organizacional", "Gestión del conocimiento", "Seminario de calidad",
				"Seminarió de gestión estrategica", "Estraregias financieras y costos de calidad"}
		}
	case 5: // Ing. en Electronica
		switch especialidad {
		case 1:
			materias = []string{"No existe una especialidad como tal"}
		}
	case 6: // Electromecanica
		switch especialidad {
		case 1:
			materias = []string{"No existe una especialidad como tal"}

		}
	case 7: // Mecatronica
		switch especialidad {
		case 1:
			materias = []string{"No existe una especialidad como tal"}
		}
	case 8: // Ing. en logistica
		switch especialidad {
		case 1:
			materias = []string{"No existe una especialidad como tal"}
		}

	default:
		fmt.Fprint(w, "Carrera no encontrada")
		return
	}

	if len(materias) == 0 {
		fmt.Fprint(w, "Especialidad no encontrada")
		return
	}

	// Crear la respuesta JSON
	respuesta := struct {
		Carrera      string   `json:"Carrera"`
		Especialidad string   `json:"Especialidad"`
		Materias     []string `json:"Materias"`
	}{
		Carrera:      nombreCarrera,
		Especialidad: obtenerNombreEspecialidad(taskID, especialidad),
		Materias:     materias,
	}

	// Escribir la respuesta JSON en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func obtenerNombreEspecialidad(ID, especialidad int) string {
	switch ID {
	case 1: // Ing. Sistemas computacionales
		switch especialidad {
		case 1:
			return "COMPUTO EMPRESARIAL"
		case 2:
			return "SISTEMAS BASADOS EN INTELIGENCIA ARTIFICIAL"
		}
	case 2: // Ing. Industrial
		switch especialidad {
		case 1:
			return "DESARROLLO DEL POTENCIAL HUMANO Y ORGANIZACIONAL"
		case 2:
			return "MANUFACTURA AVANZADA"
		}
	case 3:
		switch especialidad {
		case 1:
			return "GESTION DE SERVICIOS DE T.I EN AMBIENTES EMPRESARIALES"
		}

	case 4:
		switch especialidad {
		case 1:
			return " GESTION DE LA CALIDAD E INNOCACION DE PROCESOS"
		}

	case 5:
		switch especialidad {
		case 1:
			return "Electronica"
		}
	case 6:
		switch especialidad {
		case 1:
			return "Electromecanica"
		}
	case 7:
		switch especialidad {
		case 1:
			return "mecatronica"
		}

	case 8:
		switch especialidad {
		case 1:
			return "Logistica"
		}
	}
	return ""
}

func obtenerNombreCarreraPorID(ID int) string {
	for _, carrera := range carreras {
		if carrera.ID == ID {
			return carrera.Nombre
		}
	}
	return ""
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API Rest")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/carreras", obtenerCarreras).Methods("GET")
	router.HandleFunc("/carreras/{ID}", obtener_carrera).Methods("GET")
	router.HandleFunc("/carreras/{ID}/{numero}", obtener_especialidad).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
