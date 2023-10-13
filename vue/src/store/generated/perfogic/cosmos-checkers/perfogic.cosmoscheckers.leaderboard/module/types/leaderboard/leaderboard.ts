/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "perfogic.cosmoscheckers.leaderboard";

export interface Leaderboard {
  winners: string;
}

const baseLeaderboard: object = { winners: "" };

export const Leaderboard = {
  encode(message: Leaderboard, writer: Writer = Writer.create()): Writer {
    if (message.winners !== "") {
      writer.uint32(10).string(message.winners);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Leaderboard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLeaderboard } as Leaderboard;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.winners = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Leaderboard {
    const message = { ...baseLeaderboard } as Leaderboard;
    if (object.winners !== undefined && object.winners !== null) {
      message.winners = String(object.winners);
    } else {
      message.winners = "";
    }
    return message;
  },

  toJSON(message: Leaderboard): unknown {
    const obj: any = {};
    message.winners !== undefined && (obj.winners = message.winners);
    return obj;
  },

  fromPartial(object: DeepPartial<Leaderboard>): Leaderboard {
    const message = { ...baseLeaderboard } as Leaderboard;
    if (object.winners !== undefined && object.winners !== null) {
      message.winners = object.winners;
    } else {
      message.winners = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
