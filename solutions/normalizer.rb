#!/usr/bin/env ruby
# Solution for https://github.com/trussworks/truss-interview/blob/main/CSV_README.md

require 'csv'
require 'date'

# Collection of input before mapped normalization
raw_collection = [ ]

# To track the column names
keys = [ ]

# Put input, at this stage force the UTF8 encoding
raw_input = ARGF.read
if not raw_input.valid_encoding?
  unicoded_input = raw_input.encode(Encoding::UTF_8, invalid: :replace, undef: :replace, replace: '�')
else
  unicoded_input = raw_input
end

# This is just used to know if we are on the key line or not. Could be
# used for something else later.
line_count = 0

# Loop through line transformming into a collection of hashes
# This will make processing each one in a map more conceptual
CSV.parse(unicoded_input, encoding:'utf-8', quote_char: '"', force_quotes: true).each do |line|
  if line_count == 0
    keys = line
    line_count += 1
    next
  end

  raw_unit = { }

  keys.each_with_index do |key,index|
    raw_unit[keys[index]] = line[index]
  end

  raw_collection << raw_unit
  line_count += 1
end

# Do the bulk of normalization efforts. All UTF-8 stuff done above.
collection = raw_collection.map do |unit|
  # Write "" to all nils
  unit.each_key do |k|
    unit[k] || unit[k] = ""
  end

  # TimeStamp normalization to EST from PST and then RFC3339
  unit["Timestamp"] += " -0800"
  unit["Timestamp"] = DateTime.strptime(unit["Timestamp"], '%m/%d/%y %k:%M:%S %p %z')
  unit["Timestamp"] = unit["Timestamp"].new_offset("-0500")
  unit["Timestamp"] = DateTime.rfc3339(unit["Timestamp"].to_s)

  # Pad ZIPs
  unit["ZIP"] = unit["ZIP"].to_s.rjust(5, "0")

  # Properly capitalize names
  unit["FullName"] = unit["FullName"].capitalize

  # The FooDuration and BarDuration columns are in HH:MM:SS.MS format
  # (where MS is milliseconds); please convert them to the total number of seconds.
  foo_duration_total = 0.0
  unit["FooDuration"].split(":").each_with_index do |val,index|
    case index
    when 0
      foo_duration_total += val.to_i * 60 * 60
    when 1
      foo_duration_total += val.to_i * 60
    when 2
      foo_duration_total += val.to_f
    else
      raise StandardError.new "Unexpected Brokenness"
    end

  end

  unit["FooDuration"] = foo_duration_total

  bar_duration_total = 0.0
  unit["BarDuration"].split(":").each_with_index do |val,index|
    case index
    when 0
      bar_duration_total += val.to_i * 60 * 60
    when 1
      bar_duration_total += val.to_i * 60
    when 2
      bar_duration_total += val.to_f
    else
      raise StandardError.new "Unexpected Brokenness"
    end

  end

  unit["BarDuration"] = bar_duration_total

  # The TotalDuration column is filled with garbage data. For each row,
  # please replace the value of TotalDuration with the sum of FooDuration
  # and BarDuration.
  unit["TotalDuration"] = foo_duration_total + bar_duration_total

  unit
end

# Compile CSV for Output
final_output = CSV.generate do |csv|
  csv << keys

  collection.each do |unit|
    line = [ ]
    keys.each do |key|
      line << unit[key]
    end

    csv << line
  end
end

puts final_output

