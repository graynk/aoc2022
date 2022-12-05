const std = @import("std");
const expect = std.testing.expect;
const ArrayList = std.ArrayList;
const Allocator = std.mem.Allocator;
const test_allocator = std.testing.allocator;

const Solver = struct {
    allocator: Allocator,
    values: []ArrayList(u8) = undefined,
    instructions: [][3]u8 = undefined,

    fn parseInput(self: *Solver, input: []const u8) !void {
        var file = try std.fs.cwd().openFile(input, .{});
        defer file.close();

        var buf_reader = std.io.bufferedReader(file.reader());
        var in_stream = buf_reader.reader();

        var buf: [50]u8 = undefined;

        var list = ArrayList(ArrayList(u8)).init(self.allocator);
        defer list.deinit();
        while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            if (std.ascii.isDigit(line[1])) {
                break;
            }
            var index: usize = 1;
            while (index < line.len) : (index += 4) {
                const columnIndex = index / 4;
                if (list.items.len <= columnIndex) {
                    try list.append(ArrayList(u8).init(self.allocator));
                }
                const symbol = line[index];
                if (symbol == ' ') {
                    continue;
                }
                if (!std.ascii.isAlphabetic(symbol)) {
                    unreachable;
                }
                try list.items[columnIndex].append(symbol);
            }
        }
        _ = try in_stream.readUntilDelimiterOrEof(&buf, '\n');
        var instructions = ArrayList([3]u8).init(self.allocator);
        defer instructions.deinit();
        while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            var spliterator = std.mem.split(u8, line, " ");
            _ = spliterator.next();
            const count = try std.fmt.parseInt(u8, spliterator.next().?, 10);
            _ = spliterator.next();
            const from = try std.fmt.parseInt(u8, spliterator.next().?, 10) - 1;
            _ = spliterator.next();
            const to = try std.fmt.parseInt(u8, spliterator.next().?, 10) - 1;
            try instructions.append([_]u8{ count, from, to });
        }
        self.values = list.toOwnedSlice();
        self.instructions = instructions.toOwnedSlice();
    }

    fn cloneValues(self: *Solver) ![]ArrayList(u8) {
        var list = ArrayList(ArrayList(u8)).init(self.allocator);
        defer list.deinit();

        for (self.values) |_, index| {
            try list.append(try self.values[index].clone());
        }
        return list.toOwnedSlice();
    }

    fn partOne(self: *Solver) ![]u8 {
        var values = try self.cloneValues();
        defer self.allocator.free(values);
        defer {
            for (values) |column| {
                column.deinit();
            }
        }

        var count: u8 = 0;
        var from: u8 = 0;
        var to: u8 = 0;
        var i: usize = 0;
        for (self.instructions) |instruction| {
            count = instruction[0];
            from = instruction[1];
            to = instruction[2];

            i = 0;
            while (i < count) : (i += 1) {
                const symbol = values[from].orderedRemove(0);
                try values[to].insert(0, symbol);
            }
        }

        var fuckingString: []u8 = try self.allocator.alloc(u8, values.len);
        for (values) |column, index| {
            fuckingString[index] = column.items[0];
        }
        return fuckingString;
    }

    fn partTwo(self: *Solver) ![]u8 {
        var values = try self.cloneValues();
        defer self.allocator.free(values);
        defer {
            for (values) |column| {
                column.deinit();
            }
        }

        var forpopin = ArrayList(u8).init(self.allocator);
        defer forpopin.deinit();

        var count: u8 = 0;
        var from: u8 = 0;
        var to: u8 = 0;
        var i: usize = 0;
        for (self.instructions) |instruction| {
            count = instruction[0];
            from = instruction[1];
            to = instruction[2];

            i = 0;
            while (i < count) : (i += 1) {
                const symbol = values[from].orderedRemove(0);
                try forpopin.insert(0, symbol);
            }
            i = 0;
            while (i < count) : (i += 1) {
                const symbol = forpopin.orderedRemove(0);
                try values[to].insert(0, symbol);
            }
            forpopin.clearRetainingCapacity();
        }

        var fuckingString: []u8 = try self.allocator.alloc(u8, values.len);
        for (values) |column, index| {
            fuckingString[index] = column.items[0];
        }
        return fuckingString;
    }

    pub fn both(self: *Solver) ![2][]u8 {
        defer self.allocator.free(self.values);
        defer {
            for (self.values) |column| {
                column.deinit();
            }
        }
        defer self.allocator.free(self.instructions);
        return [2][]u8{ try self.partOne(), try self.partTwo() };
    }
};

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    var solver = Solver{ .allocator = arena.allocator() };
    try solver.parseInput("input");
    const stdout = std.io.getStdOut().writer();
    const both = try solver.both();
    defer solver.allocator.free(both[0]);
    defer solver.allocator.free(both[1]);
    try stdout.print("{s}, {s}\n", .{ both[0], both[1] });
}

test "part 1 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = try solver.both();
    defer test_allocator.free(both[0]);
    defer test_allocator.free(both[1]);
    try expect(std.mem.eql(u8, both[0], "CMZ"));
}

test "part 1 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = try solver.both();
    defer test_allocator.free(both[0]);
    defer test_allocator.free(both[1]);
    try expect(std.mem.eql(u8, both[0], "ZBDRNPMVH"));
}

test "part 2 test" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("test");
    const both = try solver.both();
    defer test_allocator.free(both[0]);
    defer test_allocator.free(both[1]);
    try expect(std.mem.eql(u8, both[1], "MCD"));
}

test "part 2 full" {
    var solver = Solver{ .allocator = test_allocator };
    try solver.parseInput("input");
    const both = try solver.both();
    defer test_allocator.free(both[0]);
    defer test_allocator.free(both[1]);
    try expect(std.mem.eql(u8, both[1], "WDLPFNNNB"));
}
