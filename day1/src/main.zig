const std = @import("std");
const expectEqual = std.testing.expectEqual;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const test_allocator = std.testing.allocator;

const Solver = struct {
    allocator: Allocator,
    values: []const u32 = undefined,
    
    fn parseInput(self: *Solver, input: []const u8) !void {
        var file = try std.fs.cwd().openFile(input, .{});
        defer file.close();

        var buf_reader = std.io.bufferedReader(file.reader());
        var in_stream = buf_reader.reader();

        var buf: [6]u8 = undefined;

        var list = ArrayList(u32).init(self.allocator);
        defer list.deinit();

        var current: u32 = 0;
        while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            if (line.len == 0) {
                try list.append(current);
                current = 0;
                continue;
            }
            const number = try std.fmt.parseInt(u32, line, 10);
            current += number;
        }
        try list.append(current);
        std.sort.sort(u32, list.items, {}, comptime std.sort.desc(u32));
        self.values = list.toOwnedSlice();
    }

    fn dayOne(self: *Solver) u32 {
        return self.values[0];
    }

    fn dayTwo(self: *Solver) u32 {
        return self.values[0] + self.values[1] + self.values[2];
    }

    pub fn both(self: *Solver) [2]u32 {
        defer self.allocator.free(self.values);
        return [2]u32{self.dayOne(), self.dayTwo()};
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
    try expectEqual(both[0], 24000);
}

test "part 1 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = solver.both();
    try expectEqual(both[0], 65912);
}

test "part 2 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = solver.both();
    try expectEqual(both[1], 45000);
}

test "part 2 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = solver.both();
    try expectEqual(both[1], 195625);
}
