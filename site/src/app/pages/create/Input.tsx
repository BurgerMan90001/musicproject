import { type JSX } from "react";
import type { ListInputState } from "../../../hooks/list";

const List = ({ list }: { list: string[] }) => {
  return (
    <ul className="scroll-vertical">
      {list.map((i) => {
        return (
          <li className="padding-xxs" key={i}>
            {i}
          </li>
        );
      })}
    </ul>
  );
};

const InputList = ({
  label,
  placeholder,
  state,
}: {
  label: string;
  placeholder: string;
  state: ListInputState;
}) => {
  return (
    <div className="display-flex flex-column">
      <div className="display-flex gap-xs">
        <label className="margin-right-auto" htmlFor={label}>
          {label}
        </label>
        <button type="button" className="button-dark" onClick={state.clear}>
          Clear
        </button>
      </div>
      <div className="bg-color-body-dark">
        <div className="display-flex">
          <input
            id={label}
            type="text"
            placeholder={placeholder}
            className="flex-1 font-size-sm padding-xs border"
            onKeyDown={(e) => {
              if (e.key != "Enter") {
                return;
              }
              e.preventDefault();
              if (state.input === "" || state.list.includes(state.input)) {
                return;
              }

              state.list.push(state.input);
              state.setInput("");
              e.currentTarget.value = "";
            }}
            onChange={(e) => {
              state.setInput(e.target.value);
            }}
          ></input>
        </div>
        <List list={state.list} />
      </div>
    </div>
  );
};
const Input = ({
  label,
  children,
}: {
  label?: string;
  children: JSX.Element;
}) => {
  return (
    <div className="">
      {label && <label htmlFor="">{label}</label>}
      <div className="display-flex bg-color-body-dark border ">{children}</div>
    </div>
  );
};
export { InputList, Input };
