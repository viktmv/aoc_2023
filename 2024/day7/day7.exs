defmodule Day7 do
  @pathname "input_test.txt"

  def run do
    File.read!(@pathname)
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(fn x -> extract(x) end)
    |> Enum.map(fn {target, operands} ->
      case solve(target, operands, [], 0) do
        true -> target
        _ -> 0
      end
    end)
    |> Enum.sum()
    |> IO.inspect()
  end

  def solve(target, [operand | tail], operators, 0) do
    solve(target, tail, operators, operand)
  end

  def solve(target, [operand | tail], operators, running_total) do
    cond do
      solve(target, tail, ["+" | operators], apply_operator("+", running_total, operand)) ->
        true

      solve(target, tail, ["*" | operators], apply_operator("*", running_total, operand)) ->
        true

      true ->
        solve(target, tail, ["||" | operators], apply_operator("||", running_total, operand))
    end
  end

  def solve(target, [], _, running_total) when running_total == target,
    do: true

  def solve(_, [], _, _), do: false

  def extract(string) do
    [target, operands] = String.split(string, ":")

    operands =
      operands
      |> String.trim()
      |> String.split(" ")
      |> Enum.map(&String.to_integer/1)

    {String.to_integer(target), operands}
  end

  def apply_operator("+", operand1, operand2), do: operand1 + operand2
  def apply_operator("*", operand1, operand2), do: operand1 * operand2

  def apply_operator("||", operand1, operand2),
    do:
      (Integer.to_string(operand1) <> Integer.to_string(operand2))
      |> String.to_integer()
end

Day7.run()
