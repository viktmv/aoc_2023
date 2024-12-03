defmodule Day2 do
  @pathname "input.txt"

  def run do
    File.read!(@pathname)
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&String.split(&1, " "))
    |> Enum.map(fn x -> Enum.map(x, &String.to_integer/1) end)
    |> Enum.map(&check_all_conditions/1)
    |> Enum.count(&(&1 == true))
    |> IO.inspect()
  end

  def check_all_conditions(list) do
    is_row_safe_asc(list) || is_row_safe_desc(list)
    # is_row_safe_asc(list, 1) || is_row_safe_desc(list, 1)
  end

  def is_row_safe_asc([_]), do: true
  def is_row_safe_asc([a, b | _]) when a == b, do: false
  def is_row_safe_asc([a, b | _]) when a > b, do: false
  def is_row_safe_asc([a, b | _]) when abs(b - a) not in 1..3, do: false
  def is_row_safe_asc([_, b | tail]), do: is_row_safe_asc([b | tail])

  def is_row_safe_desc([_]), do: true
  def is_row_safe_desc([a, b | _]) when a == b, do: false
  def is_row_safe_desc([a, b | _]) when a < b, do: false
  def is_row_safe_desc([a, b | _]) when abs(b - a) not in 1..3, do: false
  def is_row_safe_desc([_, b | tail]), do: is_row_safe_desc([b | tail])

  #  def is_row_safe_asc([_], _), do: true
  #  def is_row_safe_asc([a, b | _], 0) when a == b, do: false
  #  def is_row_safe_asc([a, b | tail], 1) when a == b, do: is_row_safe_asc([b | tail], 0)
  #  def is_row_safe_asc([a, b | _], 0) when a > b, do: false
  #  def is_row_safe_asc([a, b | tail], 1) when a > b, do: is_row_safe_asc([b | tail], 0)
  #  def is_row_safe_asc([a, b | _], 0) when abs(b - a) not in 1..3, do: false
  #
  #  def is_row_safe_asc([a, b | tail], 1) when abs(b - a) not in 1..3,
  #    do: is_row_safe_asc([b | tail], 0)
  #
  #  def is_row_safe_asc([_, b | tail], counter), do: is_row_safe_asc([b | tail], counter)
  #
  #  def is_row_safe_desc([_], _), do: true
  #  def is_row_safe_desc([a, b | _], 0) when a == b, do: false
  #  def is_row_safe_desc([a, b | tail], 1) when a == b, do: is_row_safe_desc([b | tail], 0)
  #  def is_row_safe_desc([a, b | _], 0) when a < b, do: false
  #  def is_row_safe_desc([a, b | tail], 1) when a < b, do: is_row_safe_desc([b | tail], 0)
  #  def is_row_safe_desc([a, b | _], 0) when abs(b - a) not in 1..3, do: false
  #
  #  def is_row_safe_desc([a, b | tail], 1) when abs(b - a) not in 1..3,
  #    do: is_row_safe_desc([b | tail], 0)
  #
  #  def is_row_safe_desc([_, b | tail], counter), do: is_row_safe_desc([b | tail], counter)
end

Day2.run()
