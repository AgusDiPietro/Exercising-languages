#include <stdio.h>
#include <string.h>

#define MAX_EMPLEADOS 5

// Definición de la estructura para el empleado
struct Empleado {
    char nombre[50];
    float salario;
};

// Función para ingresar datos de los empleados
void ingresarEmpleados(struct Empleado empleados[], int n) {
    for (int i = 0; i < n; ++i) {
        printf("Ingrese el nombre del empleado %d: ", i + 1);
        fgets(empleados[i].nombre, sizeof(empleados[i].nombre), stdin);
        empleados[i].nombre[strcspn(empleados[i].nombre, "\n")] ;
        printf("Ingrese el salario del empleado %d: ", i + 1);
        scanf("%f", &empleados[i].salario);
        while (getchar() != '\n'); // Limpiar el buffer de entrada
    }
}

// Función para mostrar los datos de los empleados
void mostrarEmpleados(const struct Empleado empleados[], int n) {
    printf("\nLista de empleados:\n");
    for (int i = 0; i < n; ++i) {
        printf("Empleado %d: Nombre: %s, Salario: %.2f\n", i + 1, empleados[i].nombre, empleados[i].salario);
    }
}

int main() {
    struct Empleado empleados[MAX_EMPLEADOS]; // Vector para almacenar empleados

    printf("Ingrese los datos de los empleados:\n");
    ingresarEmpleados(empleados, MAX_EMPLEADOS);

    mostrarEmpleados(empleados, MAX_EMPLEADOS);

    return 0;
}

