import { User } from "../pages/types/users";

type Props = {
  connectedUsers: User[];
};

export function ConnectedUsers({ connectedUsers }: Props) {
  const sortedUsers = connectedUsers.sort((a, b) =>
    a.name.localeCompare(b.name)
  );

  return (
    <div className="p-4 rounded bg-indigo-500/20 text-indigo-500 flex flex-col gap-2 absolute right-5 top-5 shadow hover:hidden">
      <h2 className="font-bold">Connected users</h2>
      <ul>
        {sortedUsers.map((user) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
}
