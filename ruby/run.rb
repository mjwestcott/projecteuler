#!/usr/bin/env ruby

require "json"

def run_problems
  answers = JSON.load(File.read("answers.json"))

  Dir.glob("problem*.rb").each do |filename|
    i = filename[/\d+/].to_i

    attempt = `ruby #{filename}`.to_i
    expected = answers[i.to_s]

    if attempt == expected
      puts "#{filename} ✓"
    else
      puts "** FAIL ** #{filename}: attempt=#{attempt}, expected=#{expected}"
    end
  end
end

run_problems
