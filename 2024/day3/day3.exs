defmodule Day3 do
  @pathname "input.txt"
  @pattern ~r/(do\(\))|(don't\(\))|(mul\((\d+),(\d+)\))/

  def run do
    File.read!(@pathname)
    |> scan()
    |> process()
    |> IO.inspect()
  end

  def scan(string) do
    Regex.scan(@pattern, string)
  end

  def process(_list, _capture \\ true, _sum \\ 0)
  def process([["do()" | _] | tail], _capture, sum), do: process(tail, true, sum)
  def process([["don't()" | _] | tail], _capture, sum), do: process(tail, false, sum)

  def process([[_, _, _, _, a, b] | tail], true, sum) do
    sum = sum + String.to_integer(a) * String.to_integer(b)
    process(tail, true, sum)
  end

  def process([_ | tail], false, sum), do: process(tail, false, sum)

  def process([], _capture, sum), do: sum
end

Day3.run()
