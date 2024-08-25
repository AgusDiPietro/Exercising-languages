using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp.POO
{
    /// <summary>
    /// La clase alumno sirve para modelar un alumno en el sistema...
    /// </summary>
    internal class Alumno
    {
        //Qué es una clase? Una abstracción de la realidad
        //1) ¿Qué datos definen a un alumno? Identidad
        //2) ¿Qué hace un alumno? Comportamiento

        //Atributos -> Identidad
        public string nombre;

        public int legajo;

        public DateTime fechaNacimiento;

        
        /// <summary>
        /// Método constructor por defecto 
        /// </summary>
        public Alumno()
        {
            Console.WriteLine("Constructor por defecto");
        }

        /// <summary>
        /// Método constructor sobrecargado
        /// </summary>
        /// <param name="nombre">enviar el nombre del alumno</param>
        /// <param name="legajo">enviar el legajo del alumno</param>
        /// <param name="fechaNacimiento"></param>
        public Alumno(string nombre, int legajo, DateTime fechaNacimiento) 
        {
            this.legajo = legajo;
            this.fechaNacimiento = fechaNacimiento;
        }


        //Comportamiento
        public void RendirFinal(string asignatura)
        {
            Console.WriteLine($"Usted está rindiendo la asignatura {asignatura}");
        }

        public decimal ConsultaNotaCursada(string asignatura) 
        {
            return 7;
        }

        //Con este método reescribo el ToString de la clase Object
        public override string ToString()
        {
            return $"Nombre: {this.nombre}, Legajo: {this.legajo}, Fecha Nacimiento: {this.fechaNacimiento.ToString("dd/MM/yyyy")} ";
        }
    }
}
