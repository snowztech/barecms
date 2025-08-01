import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeProvider } from "@/providers/ThemeProvider";
import { AuthProvider } from "@/providers/AuthProvider";
import { useAuth } from "@/contexts/AuthContext";
import AppSkeleton from "./components/AppSkeleton";
import Footer from "./components/Footer";
import Header from "./components/Header";
import HomePage from "./pages/Home";
import SiteDetailsPage from "@/pages/sites/Detail";
import CollectionDetailsPage from "@/pages/collections/Detail";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import PrivateRoute from "./middlewares/PrivateRoute";
import Profile from "./pages/Profile";

const AppContent = () => {
  const { initializing } = useAuth();

  if (initializing) {
    return <AppSkeleton />;
  }

  return (
    <BrowserRouter>
      <div className="min-h-screen bg-base-100 text-base-content flex flex-col">
        <Header />
        <main className="flex-1 py-8">
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route element={<PrivateRoute />}>
              <Route path="/" element={<HomePage />} />
              <Route path="/profile" element={<Profile />} />
              <Route path="/sites/:id" element={<SiteDetailsPage />} />
              <Route
                path="/sites/:siteId/collections/:id"
                element={<CollectionDetailsPage />}
              />
            </Route>
          </Routes>
        </main>
        <Footer />
      </div>
    </BrowserRouter>
  );
};

function App() {
  return (
    <ThemeProvider>
      <AuthProvider>
        <AppContent />
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;
