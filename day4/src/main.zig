const std = @import("std");
const expectEqual = std.testing.expectEqual;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const test_allocator = std.testing.allocator;

const ElfRange = [2]u8;
const ElfPair = [2]ElfRange;

const Solver = struct {
    allocator: Allocator,
    values: []ElfPair = undefined,

    fn parseInput(self: *Solver, input: []const u8) !void {
        var file = try std.fs.cwd().openFile(input, .{});
        defer file.close();

        var buf_reader = std.io.bufferedReader(file.reader());
        var in_stream = buf_reader.reader();

        var buf: [12]u8 = undefined;

        var list = ArrayList(ElfPair).init(self.allocator);
        defer list.deinit();

        while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            var pair = ElfPair{ ElfRange{ 0, 0 }, ElfRange{ 0, 0 } };

            var elfIndex: u8 = 0;
            var valueIndex: u8 = 0;
            var startIndex: usize = 0;
            const lastIndex = line.len - 1;

            // who needs .split()...
            for (line) |symbol, index| {
                const last = index == lastIndex;
                if (last) {
                    index += 1;
                }
                if (std.ascii.isDigit(symbol) and !last) {
                    continue;
                }

                pair[elfIndex][valueIndex] = try std.fmt.parseInt(u8, line[startIndex..index], 10);

                switch (symbol) {
                    '-' => {
                        valueIndex += 1;
                        startIndex = index + 1;
                    },
                    ',' => {
                        elfIndex += 1;
                        valueIndex = 0;
                        startIndex = index + 1;
                    },
                    else => break,
                }
            }
            try list.append(pair);
        }
        self.values = list.toOwnedSlice();
    }

    fn dayOne(self: *Solver) u16 {
        var count: u16 = 0;
        for (self.values) |elfPair| {
            if ((elfPair[0][0] <= elfPair[1][0] and elfPair[0][1] >= elfPair[1][1]) or (elfPair[1][0] <= elfPair[0][0] and elfPair[1][1] >= elfPair[0][1])) {
                count += 1;
            }
        }
        return count;
    }

    fn dayTwo(self: *Solver) u16 {
        var count: u16 = 0;
        for (self.values) |elfPair| {
            if (elfPair[0][0] <= elfPair[1][1] and elfPair[1][0] <= elfPair[0][1]) {
                count += 1;
            }
        }
        return count;
    }

    pub fn both(self: *Solver) [2]u16 {
        defer self.allocator.free(self.values);
        return [2]u16{ self.dayOne(), self.dayTwo() };
    }
};

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    var solver = Solver{ .allocator = arena.allocator() };
    try solver.parseInput("input");
    const stdout = std.io.getStdOut().writer();
    const both = solver.both();

    try stdout.print("{d}, {d}\n", .{ both[0], both[1] });
}

test "part 1 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = solver.both();
    try expectEqual(both[0], 2);
}

test "part 1 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = solver.both();
    try expectEqual(both[0], 498);
}

test "part 2 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = solver.both();
    try expectEqual(both[1], 4);
}

test "part 2 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = solver.both();
    try expectEqual(both[1], 859);
}
