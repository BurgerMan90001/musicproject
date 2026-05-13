import { create } from "zustand";

export interface ListState {
  list: string[];
  clear: () => void;
}

export interface ListInputState extends ListState {
  input: string;
  setInput: (v: string) => void;
}
