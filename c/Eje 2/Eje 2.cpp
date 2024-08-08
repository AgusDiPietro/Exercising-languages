#include <stdio.h>

#define MAX_ATENCIONES 100  // Define la cantidad maxima de atenciones.

typedef struct {
    int nro_atencion;
    char nombre[50];
    int servicio;
    int obra_social;
    float valor_consulta;
} Atencion;

int main() {
    Atencion atencion;  // Variable para almacenar una atención.
    int total_atenciones = 0;
    float total_cobrar_os_OSDE = 0;
    float total_cobrar_os_SwissMedical = 0;
    float total_cobrar_os_Galeno = 0;
    int total_servicios_neurologia = 0;
    int total_servicios_pediatria = 0;
    int total_servicios_nutricion = 0;
    int total_servicios_cardiologia = 0;
    int total_servicios_dermatologia = 0;
    int total_servicios_traumatologia = 0;
    int total_os_OSDE = 0;
    int total_os_SwissMedical = 0;
    int total_os_Galeno = 0;
    float valor_consultas_total = 0;
    int obra_social_mas_pacientes = 0;

    // Solicitar los datos de las atenciones
    while (total_atenciones < MAX_ATENCIONES) {
        printf("Ingrese el numero de atencion (0 para finalizar): ");
        scanf("%d", &atencion.nro_atencion);

        if (atencion.nro_atencion == 0) break;

        printf("Ingrese el nombre del paciente: ");
        scanf("%s", atencion.nombre);

        printf("Ingrese el servicio (0: Neurologia, 1: Pediatria, 2: Nutricion, 3: Cardiologia, 4: Dermatologia, 5: Traumatologia): ");
        scanf("%d", &atencion.servicio);
        
        if (atencion.servicio < 0 || atencion.servicio > 5) {
            printf("Error en la eleccion de servicio. Vuelva a cargar los datos.\n");
            break;
        }

        printf("Ingrese la obra social (0: OSDE, 1: Swiss Medical, 2: Galeno): ");
        scanf("%d", &atencion.obra_social);
        
        if (atencion.obra_social < 0 || atencion.obra_social > 2) {
            printf("Error en la eleccion de obra social. Vuelva a cargar los datos.\n");
            break;
        }

        printf("Ingrese el valor de la consulta: ");
        scanf("%f", &atencion.valor_consulta);

        // Calcular total cobrado por obra social, total de servicios y total de pacientes por obra social
        switch (atencion.obra_social) {
            case 0:
                total_cobrar_os_OSDE += atencion.valor_consulta;
                total_os_OSDE++;
                break;
            case 1:
                total_cobrar_os_SwissMedical += atencion.valor_consulta;
                total_os_SwissMedical++;
                break;
            case 2:
                total_cobrar_os_Galeno += atencion.valor_consulta;
                total_os_Galeno++;
                break;
        }

        switch (atencion.servicio) {
            case 0:
                total_servicios_neurologia++;
                break;
            case 1:
                total_servicios_pediatria++;
                break;
            case 2:
                total_servicios_nutricion++;
                break;
            case 3:
                total_servicios_cardiologia++;
                break;
            case 4:
                total_servicios_dermatologia++;
                break;
            case 5:
                total_servicios_traumatologia++;
                break;
        }

        valor_consultas_total += atencion.valor_consulta;
        total_atenciones++;
    }

    // 1. Monto total para cobrar de cada Obra Social
    printf("\nMonto total de cada obra social:\n");
    printf("OSDE: %.2f\n", total_cobrar_os_OSDE);
    printf("Swiss Medical: %.2f\n", total_cobrar_os_SwissMedical);
    printf("Galeno: %.2f\n", total_cobrar_os_Galeno);

    // 2. Valor promedio total de las consultas realizadas
    if (total_atenciones > 0) {
        float valor_promedio = valor_consultas_total / total_atenciones;
        printf("\nEl valor promedio total de las consultas realizadas es de: %.2f\n", valor_promedio);
    } else {
        printf("\nNo se realizaron consultas.\n");
    }
    
    // 3. Porcentaje de la atencion de cada servicio sobre el total
    printf("\nPorcentaje de la atención de cada servicio sobre el total:\n");
    
	float porcentaje_neurologia = (float)total_servicios_neurologia / total_atenciones * 100;
	printf("Servicio Neurología: %.2f%%\n", porcentaje_neurologia);
	
	float porcentaje_pediatria = (float)total_servicios_pediatria / total_atenciones * 100;
	printf("Servicio Pediatría: %.2f%%\n", porcentaje_pediatria);
	
	float porcentaje_nutricion = (float)total_servicios_nutricion / total_atenciones * 100;
	printf("Servicio Nutrición: %.2f%%\n", porcentaje_nutricion);
	
	float porcentaje_cardiologia = (float)total_servicios_cardiologia / total_atenciones * 100;
	printf("Servicio Cardiología: %.2f%%\n", porcentaje_cardiologia);
	
	float porcentaje_dermatologia = (float)total_servicios_dermatologia / total_atenciones * 100;
	printf("Servicio Dermatología: %.2f%%\n", porcentaje_dermatologia);

	float porcentaje_traumatologia = (float)total_servicios_traumatologia / total_atenciones * 100;
	printf("Servicio Traumatología: %.2f%%\n", porcentaje_traumatologia);

    // 4. Obra social con mas pacientes atendidos
	if (total_os_OSDE > total_os_SwissMedical && total_os_OSDE > total_os_Galeno) {
    obra_social_mas_pacientes = 1;
	} 
	else if (total_os_SwissMedical > total_os_Galeno && total_os_SwissMedical >total_os_OSDE ) {
    obra_social_mas_pacientes = 2;
	} 
	else if (total_os_Galeno > total_os_SwissMedical && total_os_Galeno >total_os_OSDE){
    obra_social_mas_pacientes = 3;
	}
	else {
		    obra_social_mas_pacientes = 0;
	}
	
	if(obra_social_mas_pacientes == 1){
		printf("\nLa obra social con mas pacientes atendidos fue la de: OSDE");
	}
	else if (obra_social_mas_pacientes == 2){
		printf("\nLa obra social con mas pacientes atendidos fue la de: SwissMedical");
	}
	else if(obra_social_mas_pacientes == 3){
		printf("\nLa obra social con mas pacientes atendidos fue la de: Galeno");
	}
	else{
        printf("\nNo se realizaron consultas.\n");
	}	
	
    // 5. Cuantas OS deben pagar mas de $100.000 al sanatorio en total
    
    int os_mayor_100k = (total_cobrar_os_OSDE > 100000) + (total_cobrar_os_SwissMedical > 100000) + (total_cobrar_os_Galeno > 100000);
    printf("\nCantidad de obras sociales que deben pagar mas de $100,000: %d\n", os_mayor_100k);

    return 0;
}

