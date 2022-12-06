const std = @import("std");
const expectEqual = std.testing.expectEqual;
const ArrayList = std.ArrayList;
const AutoArrayHashMap = std.AutoArrayHashMap;
const Allocator = std.mem.Allocator;
const test_allocator = std.testing.allocator;

const FuckingString = []u8;

const Solver = struct {
    allocator: Allocator,
    values: FuckingString = undefined,

    fn parseInput(self: *Solver, input: []const u8) !void {
        var file = try std.fs.cwd().openFile(input, .{});
        defer file.close();

        var buf_reader = std.io.bufferedReader(file.reader());
        var in_stream = buf_reader.reader();

        var buf: [4098]u8 = undefined;
        self.values = (try in_stream.readUntilDelimiterOrEof(&buf, '\n')).?;
    }

    fn partOne(self: *Solver) !usize {
        for (self.values) |_, index| {
            var i = index;
            var disOne = true;
            const n: usize = 4;
            while (i < index + n) : (i += 1) {
                var j = i + 1;
                while (j < index + n) : (j += 1) {
                    if (i == j) {
                        continue;
                    }
                    if (self.values[i] == self.values[j]) {
                        disOne = false;
                    }
                }
            }
            if (disOne) {
                return index + n;
            }
        }
        return error.InvalidParam;
    }

    fn partTwo(self: *Solver) !usize {
        for (self.values) |_, index| {
            var i = index;
            var disOne = true;
            const n: usize = 14;
            while (i < index + n) : (i += 1) {
                var j = i + 1;
                while (j < index + n) : (j += 1) {
                    if (i == j) {
                        continue;
                    }
                    if (self.values[i] == self.values[j]) {
                        disOne = false;
                    }
                }
            }
            if (disOne) {
                return index + n;
            }
        }
        return error.InvalidParam;
    }

    pub fn both(self: *Solver) ![2]usize {
        return [2]usize{ try self.partOne(), try self.partTwo() };
    }
};

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    var solver = Solver{ .allocator = arena.allocator() };
    try solver.parseInput("input");
    const stdout = std.io.getStdOut().writer();
    const both = try solver.both();

    try stdout.print("{d}, {d}\n", .{ both[0], both[1] });
}

test "part 1 test 1" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = try solver.both();
    try expectEqual(both[0], 5);
}

test "part 1 test 2" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test2");
    const both = try solver.both();
    try expectEqual(both[0], 6);
}

test "part 1 test 3" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test3");
    const both = try solver.both();
    try expectEqual(both[0], 10);
}

test "part 1 test 4" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test4");
    const both = try solver.both();
    try expectEqual(both[0], 11);
}

test "part 1 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = try solver.both();
    try expectEqual(both[0], 1544);
}

test "part 2 test 1" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = try solver.both();
    try expectEqual(both[1], 23);
}

test "part 2 test 2" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test2");
    const both = try solver.both();
    try expectEqual(both[1], 23);
}

test "part 2 test 3" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test3");
    const both = try solver.both();
    try expectEqual(both[1], 29);
}

test "part 2 test 4" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test4");
    const both = try solver.both();
    try expectEqual(both[1], 26);
}

test "part 2 test 5" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test5");
    const both = try solver.both();
    try expectEqual(both[1], 19);
}

test "part 2 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = try solver.both();
    try expectEqual(both[1], 2145);
}
