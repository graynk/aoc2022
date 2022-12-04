const std = @import("std");
const expectEqual = std.testing.expectEqual;
const ArrayList = std.ArrayList;
const AutoArrayHashMap = std.AutoArrayHashMap;
const Allocator = std.mem.Allocator;
const test_allocator = std.testing.allocator;

const FuckingString = []u8;

const Solver = struct {
    allocator: Allocator,
    values: []FuckingString = undefined,

    fn parseInput(self: *Solver, input: []const u8) !void {
        var file = try std.fs.cwd().openFile(input, .{});
        defer file.close();

        var buf_reader = std.io.bufferedReader(file.reader());
        var in_stream = buf_reader.reader();

        var buf: [1024]u8 = undefined;

        var list = ArrayList(FuckingString).init(self.allocator);
        defer list.deinit();

        while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            try list.append(try self.allocator.dupe(u8, line));
        }
        self.values = list.toOwnedSlice();
    }

    fn lineToSet(set: *AutoArrayHashMap(u8, void), line: []u8) !void {
        for (line) |symbol| {
            try set.*.put(symbol, {});
        }
    }

    fn symbolToPriority(symbol: u8) u16 {
        return switch (symbol) {
            'A'...'Z' => symbol - 'A' + 27,
            'a'...'z' => symbol - 'a' + 1,
            else => unreachable,
        };
    }

    fn intersect(_: *Solver, sets: []const AutoArrayHashMap(u8, void)) !u8 {
        for (sets) |set| {
            var it = set.iterator();
            while (it.next()) |entry| {
                var hasCommon = true;
                const symbol = entry.key_ptr.*;

                for (sets) |otherSet| {
                    if (!otherSet.contains(symbol)) {
                        hasCommon = false;
                        break;
                    }
                }
                if (hasCommon) {
                    return symbol;
                }
            }
        }

        return error.InvalidParam;
    }

    fn partOne(self: *Solver) !u16 {
        var list = ArrayList(u8).init(self.allocator);
        defer list.deinit();
        var sets = [2]AutoArrayHashMap(u8, void){ AutoArrayHashMap(u8, void).init(self.allocator), AutoArrayHashMap(u8, void).init(self.allocator) };
        defer {
            sets[0].deinit();
            sets[1].deinit();
        }
        for (self.values) |line| {
            const middle = line.len / 2;
            try lineToSet(&sets[0], line[0..middle]);
            try lineToSet(&sets[1], line[middle..]);
            const commonSymbol = try self.intersect(&sets);
            try list.append(commonSymbol);
            sets[0].clearRetainingCapacity();
            sets[1].clearRetainingCapacity();
        }
        var priorities: u16 = 0;
        for (list.items) |symbol| {
            priorities += symbolToPriority(symbol);
        }
        return priorities;
    }

    fn partTwo(self: *Solver) !u16 {
        var priorities: u16 = 0;

        var sets = [3]AutoArrayHashMap(u8, void){ AutoArrayHashMap(u8, void).init(self.allocator), AutoArrayHashMap(u8, void).init(self.allocator), AutoArrayHashMap(u8, void).init(self.allocator) };
        defer {
            var j: usize = 0;
            while (j < 3) : (j += 1) {
                sets[j].deinit();
            }
        }
        var i: usize = 0;
        while (i < self.values.len) : (i += 3) {
            var j: usize = 0;
            while (j < 3) : (j += 1) {
                try lineToSet(&sets[j], self.values[j + i]);
            }
            priorities += symbolToPriority(try self.intersect(&sets));
            j = 0;
            while (j < 3) : (j += 1) {
                sets[j].clearRetainingCapacity();
            }
        }

        return priorities;
    }

    pub fn both(self: *Solver) ![2]u16 {
        defer self.allocator.free(self.values);
        defer {
            for (self.values) |value| {
                self.allocator.free(value);
            }
        }
        return [2]u16{ try self.partOne(), try self.partTwo() };
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

test "part 1 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = try solver.both();
    try expectEqual(both[0], 157);
}

test "part 1 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = try solver.both();
    try expectEqual(both[0], 8153);
}

test "part 2 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = try solver.both();
    try expectEqual(both[1], 70);
}

test "part 2 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = try solver.both();
    try expectEqual(both[1], 2342);
}
