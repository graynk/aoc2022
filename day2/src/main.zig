const std = @import("std");
const expectEqual = std.testing.expectEqual;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const test_allocator = std.testing.allocator;

const Round = [2]u8; // 0 - Opponent, 1 - Me

const Solver = struct {
    allocator: Allocator,
    values: []Round = undefined,

    fn parseInput(self: *Solver, input: []const u8) !void {
        var file = try std.fs.cwd().openFile(input, .{});
        defer file.close();

        var buf_reader = std.io.bufferedReader(file.reader());
        var in_stream = buf_reader.reader();

        var buf: [4]u8 = undefined;

        var list = ArrayList(Round).init(self.allocator);
        defer list.deinit();

        while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            const round = Round{ line[0], line[2] };
            try list.append(round);
        }
        self.values = list.toOwnedSlice();
    }

    fn partOne(self: *Solver) u16 {
        var count: u16 = 0;
        for (self.values) |round| {
            count += round[1] - 'X' + 1;
            switch (round[1] - round[0]) {
                21, 24 => count += 6,
                23 => count += 3,
                else => count += 0,
            }
        }
        return count;
    }

    fn partTwo(self: *Solver) u16 {
        var count: u16 = 0;
        for (self.values) |round| {
            switch (round[1]) {
                'X' => {
                    var res: u8 = round[0] - 'A';
                    if (res == 0) {
                        res = 3;
                    }
                    count += res;
                },
                'Y' => count += 3 + round[0] - 'A' + 1,
                'Z' => count += 6 + (round[0] - 'A' + 1) % 3 + 1,
                else => unreachable,
            }
        }
        return count;
    }

    pub fn both(self: *Solver) [2]u16 {
        defer self.allocator.free(self.values);
        return [2]u16{ self.partOne(), self.partTwo() };
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
    try expectEqual(both[0], 15);
}

test "part 1 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = solver.both();
    try expectEqual(both[0], 15422);
}

test "part 2 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = solver.both();
    try expectEqual(both[1], 12);
}

test "part 2 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = solver.both();
    try expectEqual(both[1], 15442);
}
