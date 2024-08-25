using ConsoleApp.POO;
using System;
using System.Collections.Generic;
using System.ComponentModel.Design.Serialization;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp
{
    internal class Program
    {
        private static int Sumar(int a, int b)
        {
            return a + b;
        }
        static void Main(string[] args)
        {
            //Introducción rápida a orientación a objetos
            Alumno alumno1;
            alumno1 = new Alumno("Gastoncito", 1234, DateTime.Now.AddDays(-40)); //Instanciar es cuando creo el objeto con new
            //Declaro un objeto llamado alumno1 del tipo Alumno
            //Luego, lo instancio utilizando la palabra reservada new.
            //alumno1.legajo = 9999;
            //alumno1.fechaNacimiento = DateTime.Now.AddYears(-20);
            //alumno1.nombre = "Pedrito";
            Console.WriteLine(alumno1);

            Alumno alumno2 = new Alumno();
            alumno2.legajo = 11;
            alumno2.fechaNacimiento = DateTime.Now.AddYears(-35);
            alumno2.nombre = "Juancito";
            Console.WriteLine(alumno2);


            Console.WriteLine("¿Cómo sé que alumno1 y alumno2 son dos objetos diferentes?");
            Console.WriteLine(alumno1 == alumno2); //Con esto comparo las referencias en memoria (Puntero)

            Alumno[] alumnos = new Alumno[2];
            alumnos[0] = alumno1;
            alumnos[1] = alumno2;

            for (int i = 0; i < alumnos.Length; i++)
            {
                Console.WriteLine("Invocamos al alumno:");
                Console.WriteLine(alumnos[i]);

                alumnos[i].RendirFinal("Programación 1");
                Console.WriteLine($"Nota cursada {alumnos[i].ConsultaNotaCursada("Programación estructurada")}");
            }

            //RepasoEstructurada();

            Console.Read();
        }

        private static void RepasoEstructurada()
        {
            int total = Sumar(4, 5);

            Console.WriteLine("Bienvenidos a C#");

            //Comentarios igual en C y /* */
            //Estructurada

            int a = 20;
            int b = 5;

            if (a > b)
            {
                Console.WriteLine("A: " + a + " es mayor que b");

                //Interpolación de String -> C# 7 en adelante
                Console.WriteLine($"A: {a} es mayor que B: {b}");
            } //else if y else


            int categoriaIngreso = 0;

            Console.WriteLine("Por favor ingrese una categoría");
            categoriaIngreso = int.Parse(Console.ReadLine());

            EnumCategoria enumCategoria = (EnumCategoria)categoriaIngreso;

            switch (enumCategoria)
            {
                case EnumCategoria.OSDE:
                    Console.WriteLine("Usted ingresó categoría OSDE");
                    break;
                case EnumCategoria.GALENO:
                    Console.WriteLine("Usted ingresó categoría GALENO");
                    break;
                case EnumCategoria.SWISS:
                    Console.WriteLine("Usted ingresó categoría SWISS");
                    break;
                default:
                    Console.WriteLine("Categoría inválida");
                    break;
            }

            //Ciclos

            for (int i = 0; i < 10; i++)
            {
                Console.WriteLine($"Estamos en el paso: {i}");
            }

            while (a > 0)
            {
                Console.WriteLine("Pasando por acá");
                a--;
            }
            do
            {
                Console.WriteLine("Es un ciclo que ejecuta al menos una vez");
                b--;
            } while (b > 0);

            int[] ints = new int[10];
            int lectura = 0;

            for (int i = 0; i < 10; i++)
            {
                Console.WriteLine("Por favor ingrese un nro");
                lectura = int.Parse(Console.ReadLine());
                ints[i] = lectura;
            }
        }
    }
}
