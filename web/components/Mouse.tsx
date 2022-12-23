import { User } from "../pages/types/users";

type Props = {
  user: User;
};

export function Mouse({ user }: Props) {
  const style = {
    left: user.mouseX,
    top: user.mouseY,
  };

  return (
    <span
      style={style}
      className="absolute px-3 py-0.5 rounded 
                 bg-indigo-500/20 pointer-events-none
                 text-indigo-500 text-xs font-semibold"
    >
      {user.name}
    </span>
  );
}
