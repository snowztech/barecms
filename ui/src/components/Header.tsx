import { useAuth } from "@/hooks/useAuth";
import { useUser } from "@/hooks/useUser";

const Header = () => {
  const { logout } = useAuth();
  const { user } = useUser();

  return (
    <header className="container-bare flex justify-between items-center py-6">
      <a
        href="/"
        className="logo flex items-center gap-2 cursor-pointer hover:opacity-80 transition-opacity"
      >
        <img src="/icon.svg" alt="logo" className="w-7 h-7" />
        <span className="text-display text-2xl">barecms.</span>
      </a>
      <div className="actions">
        {user && (
          <div className="dropdown dropdown-end">
            <div
              tabIndex={0}
              role="button"
              className="font-medium text-base-content hover:text-primary transition-colors cursor-pointer focus-bare"
            >
              Hello, {user.username}
            </div>
            <ul
              tabIndex={0}
              className="dropdown-content dropdown-bare menu bg-base-100 rounded-bare z-[1] p-2 w-40 shadow-bare-lg mt-2"
            >
              <li>
                <a
                  href="/profile"
                  className="text-sm font-medium hover:bg-base-200 rounded transition-colors"
                >
                  Profile
                </a>
              </li>
              <li>
                <button
                  onClick={logout}
                  className="text-sm font-medium hover:bg-base-200 rounded transition-colors text-left w-full"
                >
                  Logout
                </button>
              </li>
            </ul>
          </div>
        )}
      </div>
    </header>
  );
};

export default Header;
