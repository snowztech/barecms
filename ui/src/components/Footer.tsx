import { Github } from "lucide-react";

const Footer = () => {
  return (
    <footer className="border-t border-bare-200 py-8 mt-16">
      <section className="container-bare flex flex-col md:flex-row gap-4 justify-between items-center">
        <span className="text-sm text-bare-600">
          © {new Date().getFullYear()} BareCMS
        </span>
        <a
          href="https://github.com/lucasnevespereira/barecms"
          target="_blank"
          rel="noopener noreferrer"
          className="flex items-center gap-2 text-sm text-bare-600 hover:text-primary transition-colors"
        >
          <Github size={16} />
          <span>View on GitHub</span>
        </a>
      </section>
    </footer>
  );
};

export default Footer;
