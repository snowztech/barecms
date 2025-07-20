import { useAuth } from "@/hooks/useAuth";
import { useUser } from "@/hooks/useUser";
import ThemeToggle from "./ThemeToggle";

const Header = () => {
  const { logout } = useAuth();
  const { user } = useUser();

  return (
    <header className="border-b border-bare-200">
      <div className="container-bare flex justify-between items-center py-6">
        <a
          href="/"
          className="logo flex items-center gap-2 cursor-pointer hover:opacity-80 transition-opacity"
        >
          <img src="/logo.png" alt="logo" className="w-7 h-7" />
          <span className="text-display text-2xl">barecms.</span>
        </a>

        <div className="flex items-center gap-4">
          <ThemeToggle />
          {user && (
            <div className="dropdown dropdown-end">
              <div
                tabIndex={0}
                role="button"
                className="font-medium text-base-content hover:text-primary transition-colors cursor-pointer focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/20 focus-visible:ring-offset-2"
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
      </div>
    </header>
  );
};

export default Header;
