using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ProjectEuler
{
	class Program
	{
		static void Main(string[] args)
		{
			bool running = true;

			while (running)
			{
				Console.WriteLine("Pick A Solution to Run:");
				Console.WriteLine("1 - Multiples of 3 and 5");
				Console.WriteLine("2 - Even Fibonacci Number");
				Console.WriteLine("3 - Largest Prime Factor");
				Console.WriteLine("-1 - Close");

				int i = int.Parse(Console.ReadLine());
				switch (i)
				{
					case 1:
						Console.WriteLine(MuliplesOf3And5());
						break;
					case 2:
						Console.WriteLine(EvenFibonacciNumbers());
						break;
					case 3:
						Console.WriteLine(LargestPrimeFactor(600851475143));
						break;
					case -1:
						running = false;
						break;
				}
				Console.ReadLine();
			}
		}

		static public int MuliplesOf3And5()
		{
			int sum = 0;

			for (int i = 3; i < 1000; i++)
				if (isMultiple(i, 3) || isMultiple(i, 5))
					sum += i;
			
			return sum;
		}

		static private bool isMultiple(long a, long b)
		{
			if (a % b == 0)
				return true;

			return false;
		}

		static public int EvenFibonacciNumbers()
		{
			int sum = 2;

			for (int a = 1, b = 2, temp; a + b < 4000000; temp = a + b, a = b, b = temp)
				if (isMultiple(a + b, 2))
					sum += a + b;

			return sum;
		}

		static public long LargestPrimeFactor(long num)
		{
			for (long i = num; i >= 1; i -= 2)
				if (isMultiple(num, i) && isPrime(i))
					return i;

			return -1;
		}

		static private bool isPrime(long a) {
			for (int i = 2; i < a / 2; i++)
				if (a % i == 0)
					return false;

			return true;
		}
	}
}
