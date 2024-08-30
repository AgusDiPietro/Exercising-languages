using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tp_inicial
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string cadena;
            int a, b, c;
            Console.WriteLine("ingrese el valor de la hora");
            cadena = Console.ReadLine();
            a = Convert.ToInt32(cadena);
            Console.WriteLine("ingrese el tiempo trabajado");
            cadena = Console.ReadLine();
            b = Convert.ToInt32(cadena);
            c = a * b;
            Console.WriteLine(" el sueldo es {0} ", c);
            Console.ReadKey();
        }
    }
}
