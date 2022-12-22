import { User } from "../pages/types/users";

type Props = {
  connectedUsers: User[];
};

export function ConnectedUsers({ connectedUsers }: Props) {
  return (
    <div className="p-4 rounded bg-blue-100 flex flex-col gap-2 absolute right-5 top-5 shadow hover:hidden">
      <h2 className="font-bold">Connected users</h2>
      <ul>
        {connectedUsers.map((user) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
}