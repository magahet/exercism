object SumOfMultiples {
    def sum(factors: Set[Int], limit: Int): Int = {
        factors.flatMap(f => f to limit by f).sum
    }
}

