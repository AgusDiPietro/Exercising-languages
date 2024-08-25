#include <stdio.h>

#define NUM_REMEDIOS 3
#define NUM_SUCURSALES 3

typedef struct {
    int nRemedio;
    float precio;
} Remedio;

typedef struct {
    int nPedido;
    int nSucursal;
    int nRemedio;
    int cantidad;
} Venta;

void ingresarRemedios(Remedio remedios[]) {
    for (int i = 0; i < NUM_REMEDIOS; i++) {
        printf("Ingrese el precio del remedio %d: ", i + 1);
        remedios[i].nRemedio = i + 1;
        scanf("%f", &remedios[i].precio);
    }
}

void procesarVentas(Venta ventas[], int *numVentas) {
    printf("Ingrese las ventas (n de pedido 0 para terminar):\n");
    for (*numVentas = 0; ; (*numVentas)++) {
        printf("N de pedido: ");
        scanf("%d", &ventas[*numVentas].nPedido);
        if (ventas[*numVentas].nPedido == 0) break;

        do {
            printf("N de sucursal (1-3): ");
            scanf("%d", &ventas[*numVentas].nSucursal);
            if (ventas[*numVentas].nSucursal < 1 || ventas[*numVentas].nSucursal > NUM_SUCURSALES) {
                printf("Error: num. de sucursal invalido. Debe ser entre 1 y 3.\n");
            }
        } while (ventas[*numVentas].nSucursal < 1 || ventas[*numVentas].nSucursal > NUM_SUCURSALES);

        do {
            printf("N de remedio (1-3): ");
            scanf("%d", &ventas[*numVentas].nRemedio);
            if (ventas[*numVentas].nRemedio < 1 || ventas[*numVentas].nRemedio > NUM_REMEDIOS) {
                printf("Error: num. de remedio invalido. Debe ser entre 1 y 3.\n");
            }
        } while (ventas[*numVentas].nRemedio < 1 || ventas[*numVentas].nRemedio > NUM_REMEDIOS);

        printf("Cantidad: ");
        scanf("%d", &ventas[*numVentas].cantidad);
    }
}

void calcularRecaudaciones(Venta ventas[], int numVentas, Remedio remedios[], float recaudacion[NUM_SUCURSALES][NUM_REMEDIOS], int cantidadVendida[NUM_SUCURSALES][NUM_REMEDIOS]) {
    for (int i = 0; i < numVentas; i++) {
        int sucursal = ventas[i].nSucursal - 1;
        int remedio = ventas[i].nRemedio - 1;
        float totalVenta = ventas[i].cantidad * remedios[remedio].precio;
        if (ventas[i].nSucursal == 1) {
            totalVenta *= 0.95;
        }
        recaudacion[sucursal][remedio] += totalVenta;
        cantidadVendida[sucursal][remedio] += ventas[i].cantidad;
    }
}

void imprimirResultados(float recaudacion[NUM_SUCURSALES][NUM_REMEDIOS], int cantidadVendida[NUM_SUCURSALES][NUM_REMEDIOS]) {
    
	// 1:Cantidad recaudada por cada remedio en cada sucursal, ordenados de menor a mayor
    for (int i = 0; i < NUM_SUCURSALES; i++) {
        printf("Sucursal %d:\n", i + 1);
        for (int j = 0; j < NUM_REMEDIOS; j++) {
            printf("  Remedio %d: %.2f\n", j + 1, recaudacion[i][j]);
        }
    }
     printf("\n");

    // 2:Remedio que se vendio menos en cada sucursal
    for (int i = 0; i < NUM_SUCURSALES; i++) {
        int menorRemedio = 0;
        for (int j = 1; j < NUM_REMEDIOS; j++) {
            if (cantidadVendida[i][j] < cantidadVendida[i][menorRemedio]) {
                menorRemedio = j;
            }
        }
        printf("La sucursal %d vendio menos del remedio %d\n", i + 1, menorRemedio + 1);
    }
     printf("\n");

    // 3:Sucursal donde se recaudó más en total
    int mayorSucursal = 0;
    float totalRecaudado[NUM_SUCURSALES] = {0};
    for (int i = 0; i < NUM_SUCURSALES; i++) {
        for (int j = 0; j < NUM_REMEDIOS; j++) {
            totalRecaudado[i] += recaudacion[i][j];
        }
        if (totalRecaudado[i] > totalRecaudado[mayorSucursal]) {
            mayorSucursal = i;
        }
    }
    printf("La sucursal que recaudo mas fue la sucursal %d\n", mayorSucursal + 1);
}

void pedidoMayorValor(Venta ventas[], int numVentas, Remedio remedios[]) {
    int nPedidoMayor = 0;
    float valorMayor = 0;
    for (int i = 0; i < numVentas; i++) {
        float valorActual = ventas[i].cantidad * remedios[ventas[i].nRemedio - 1].precio;
        if (ventas[i].nSucursal == 1) {
            valorActual *= 0.95;
        }
        if (valorActual > valorMayor) {
            valorMayor = valorActual;
            nPedidoMayor = ventas[i].nPedido;
        }
    }
    printf("El pedido de mayor valor fue el : %d\n", nPedidoMayor);
     printf("\n");
}

 
void porcentajeRecaudacion(float recaudacion[NUM_SUCURSALES][NUM_REMEDIOS]) {
    float totalGlobal = 0;
    float totalRecaudado[NUM_SUCURSALES] = {0};
    for (int i = 0; i < NUM_SUCURSALES; i++) {
        for (int j = 0; j < NUM_REMEDIOS; j++) {
            totalRecaudado[i] += recaudacion[i][j];
        }
        totalGlobal += totalRecaudado[i];
    }
    for (int i = 0; i < NUM_SUCURSALES; i++) {
        printf("Sucursal %d: %.2f%%\n", i + 1, (totalRecaudado[i] / totalGlobal) * 100);
    }
}

int main() {
    Remedio remedios[NUM_REMEDIOS];
    Venta ventas[50]; // Suponemos que no habrán más de 50 ventas
    int numVentas;
    float recaudacion[NUM_SUCURSALES][NUM_REMEDIOS] = {0};
    int cantidadVendida[NUM_SUCURSALES][NUM_REMEDIOS] = {0};

    ingresarRemedios(remedios);
    procesarVentas(ventas, &numVentas);
    calcularRecaudaciones(ventas, numVentas, remedios, recaudacion, cantidadVendida);
    imprimirResultados(recaudacion, cantidadVendida);
    pedidoMayorValor(ventas, numVentas, remedios);
    porcentajeRecaudacion(recaudacion);

    return 0;
}

