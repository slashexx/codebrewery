import React from 'react';
import { Code2, Sparkles, Terminal, Cpu, Beer, Wine } from 'lucide-react';
import { useNavigate } from 'react-router-dom';

export default function IntroPage() {
    const navigate = useNavigate();

  return (
    <div className="min-h-screen bg-black relative overflow-hidden flex items-center justify-center">
      {/* Animated background elements */}
      <div className="absolute inset-0 overflow-hidden">
        <div className="absolute -left-1/4 -top-1/4 w-1/2 h-1/2 bg-blue-500/20 rounded-full blur-[120px] animate-pulse"></div>
        <div className="absolute -right-1/4 -bottom-1/4 w-1/2 h-1/2 bg-purple-500/20 rounded-full blur-[120px] animate-pulse delay-700"></div>
        <div className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 w-1/3 h-1/3 bg-cyan-500/20 rounded-full blur-[100px] animate-pulse delay-1000"></div>
      </div>

      {/* Floating code symbols */}
      <div className="absolute inset-0 overflow-hidden">
        <div className="absolute top-[15%] left-[10%] animate-float-slow">
          <Terminal className="w-8 h-8 text-blue-400/40" />
        </div>
        <div className="absolute top-[45%] right-[15%] animate-float-slower">
          <Cpu className="w-10 h-10 text-purple-400/40" />
        </div>
        <div className="absolute bottom-[20%] left-[20%] animate-float">
          <Code2 className="w-12 h-12 text-cyan-400/40" />
        </div>
      </div>

      {/* Main content */}
      <div className="relative z-10 max-w-4xl mx-auto px-4 text-center backdrop-blur-lg  rounded-xl p-8">
        {/* Logo section */}
        <div className="mb-8 inline-flex items-center gap-4 backdrop-blur-sm bg-white/5 px-6 py-2 rounded-2xl border border-white/10">
          <Wine className="w-8 h-8 text-blue-400" />
          <h1 className="text-3xl font-bold bg-gradient-to-r from-blue-400 via-cyan-400 to-purple-400 text-transparent bg-clip-text">
            CodeBrewery
          </h1>
        </div>

        {/* Main heading */}
            <Sparkles className="absolute -top-8 -right-8 w-6 h-6 text-yellow-400 animate-pulse" />
        <h2 className="text-5xl md:text-7xl font-bold text-white mb-6 leading-tight">
          Run code. Anytime.
          <span className="relative inline-block ml-4">
            Anywhere.
          </span>
        </h2>

        {/* Description */}
        <p className="text-lg md:text-xl text-gray-300 mb-12 max-w-2xl mx-auto">
          Experience the future of coding with our powerful online IDE. Write, compile, and execute code in multiple languages with zero setup.
        </p>

        {/* CTA Button */}
        <div className="flex justify-center gap-6">
          <button 
            onClick={() => {navigate('/editor')}} 
            className="group relative inline-flex items-center gap-2 px-8 py-4 bg-gradient-to-r from-blue-600 to-cyan-600 rounded-xl text-white font-semibold text-lg transition-all duration-300 hover:scale-105 hover:shadow-[0_0_40px_rgba(59,130,246,0.3)] focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-gray-900"
          >
            <span>Code Now!</span>
            <Code2 className="w-5 h-5 transition-transform group-hover:translate-x-1" />
          </button>
        </div>

        {/* Tagline */}
        {/* <p className="absolute bottom-8 left-1/2 -translate-x-1/2 text-gray-400 font-light tracking-wider">
          Run code. Anytime. Anywhere.
        </p> */}
      </div>

      {/* Footer */}
      <footer className="absolute bottom-4 left-1/2 -translate-x-1/2 text-gray-400 text-sm">
        Made with <span className="text-red-500">♥</span> by slashexx
      </footer>
    </div>
  );
}
