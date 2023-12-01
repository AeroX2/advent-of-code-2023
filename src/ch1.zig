const std = @import("std");
const regex = @import("regex/src/regex.zig");

const allocator = std.heap.page_allocator;
const ArrayList = std.ArrayList;
const print = std.debug.print;

const Regex = regex.Regex;

pub fn main() !void {
    var args = std.process.args();
    _ = args.skip();

    const input_file = args.next() orelse "";
    std.debug.print("Running ch1 with {s}\n", .{input_file});

    const file = try read_file(input_file);
    const sum = part1(file);
    print("Part 1 Sum {}\n", .{sum});

    const sum2 = part2(file);
    print("Part 2 Sum {}\n", .{sum2});

    defer allocator.free(file);
}

pub fn read_file(input_file: []const u8) ![]u8 {
    var file = try std.fs.cwd().readFileAlloc(allocator, input_file, 100000);
    return file;
}

pub fn part1(file: []u8) i32 {
    var lines = std.mem.tokenize(u8, file, "\n");

    var sum: i32 = 0;
    while (lines.next()) |line| {
        // std.debug.print("{s}\n", .{line});

        var a: i32 = 0;
        for (line) |c| {
            if (std.ascii.isDigit(c)) {
                a += c - '0';
                break;
            }
        }

        var b: i32 = 0;
        var i: usize = line.len - 1;
        while (i >= 0) : (i -= 1) {
            var c = line[i];
            if (std.ascii.isDigit(c)) {
                b += c - '0';
                break;
            }
        }
        sum += a * 10 + b;
    }
    return sum;
}

const map = std.ComptimeStringMap(i32, .{
    .{ "one", 1 },
    .{ "two", 2 },
    .{ "three", 3 },
    .{ "four", 4 },
    .{ "five", 5 },
    .{ "six", 6 },
    .{ "seven", 7 },
    .{ "eight", 8 },
    .{ "nine", 9 },
    .{ "zero", 0 },
});

const numbers = "one|two|three|four|five|six|seven|eight|nine|zero";
pub fn part2(file: []u8) i32 {
    const re = try Regex.compile(allocator, "(" ++ numbers ++ "|[0-9]).*(" ++ numbers ++ "|[0-9])");
    const captures = try re.captures(file);
    print("{s} {s}", .{ captures.sliceAt(1), captures.sliceAt(1) });
    return 0;
}
