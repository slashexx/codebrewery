import React, { useState } from "react";
import Editor from "@monaco-editor/react";
import {
  Code2,
  Play,
  Cpu,
  Terminal,
  Wand2,
  RotateCcw,
  Wine,
  Coffee
} from "lucide-react";
import { useNavigate } from "react-router-dom";
import "devicon/devicon.min.css";


type Language = "c" | "cpp" | "python" | "java" | "go" | "rs";

const INITIAL_CODE: Record<Language, string> = {
  c: '#include <stdio.h> \n\nint main() {\n    printf("Hello world !"); \n    return 0; \n}',
  cpp: '#include <iostream>\n\nint main() {\n    std::cout << "Hello, World!" << std::endl;\n    return 0;\n}',
  python: 'print("Hello, World!")',
  java: 'public class Main {\n    public static void main(String[] args) {\n        System.out.println("Hello, World!");\n    }\n}',
  go: 'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}',
  rs: 'fn main() {\n\tprintln!("Hello world !")\n}'
};

const LANGUAGE_CONFIGS = {
  c: { label: 'C', icon: "devicon-c-plain colored", color: "text-red-500" },
  cpp: { label: "C++", icon:"devicon-cplusplus-plain colored", color: "text-blue-500" },
  python: { label: "Python", icon: "devicon-python-plain colored" , color: "text-yellow-500" },
  java: { label: "Java", icon: "devicon-java-plain colored", color: "text-red-500" },
  go: { label: "Go", icon: "devicon-go-plain colored", color: "text-cyan-500" },
  rs: { label: "Rust", icon: "devicon-rust-original colored", color: "text-red-500"}
};

function App() {
  const [selectedLanguage, setSelectedLanguage] = useState<Language>("c");
  const [code, setCode] = useState(INITIAL_CODE[selectedLanguage]);
  const [output, setOutput] = useState("Output will appear here...");
  const [editorInstance, setEditorInstance] = useState<monaco.editor.IStandaloneCodeEditor>(null);
  const navigate = useNavigate();

  const handleLanguageChange = (language: Language) => {
    setSelectedLanguage(language);
    setCode(INITIAL_CODE[language]);
  };

  const handleRunCode = async () => {
    setOutput(`Running ${LANGUAGE_CONFIGS[selectedLanguage].label} code...\n`);

    try {
      const response = await fetch("https://codebrewery-api-gateway.onrender.com/execute", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ language: selectedLanguage, code }),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const result = await response.json();
      console.log("Execution Result:", result);
      // Display the result in your UI
      if (result.error == null) {
        setOutput(`${result.output}\n`);
      } else {
        setOutput(`${result.error}\n ${result.output}\n`);
      }
    } catch (error) {
      console.error("Error executing code:", error);
    }
  };

  const handleFormatCode = () => {
    if (editorInstance) {
      editorInstance.getAction("editor.action.formatDocument").run();
    }
  };

  const handleResetCode = () => {
    setCode(INITIAL_CODE[selectedLanguage]);
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700 p-4">
        <div className="max-w-[1920px] mx-auto flex justify-between items-center">
          <div className="flex items-center gap-2 cursor-pointer">
            <Wine className="w-8 h-8 text-blue-400" />
            <h1
              className="text-2xl font-bold bg-gradient-to-r from-blue-400 via-cyan-400 to-purple-400 text-transparent bg-clip-text"
              onClick={() => {
                navigate("/");
              }}
            >
              CodeBrewery
            </h1>
          </div>
          <button
            onClick={handleRunCode}
            className="relative flex items-center gap-2 px-6 py-2 rounded-lg overflow-hidden group"
          >
            <div className="absolute inset-0 bg-gradient-to-r from-blue-400 via-cyan-400 to-purple-400 opacity-80 rounded-lg animate-gradient"></div>
            <span className="relative z-10 text-white font-semibold">Run</span>
            <Play className="relative z-10 w-4 h-4" />
          </button>
        </div>
      </header>

      <div className="flex h-[calc(100vh-4rem)]">
        {/* Language Selection Sidebar */}
        <div className="w-20 bg-gray-800 border-r border-gray-700 p-2 flex flex-col gap-2">
          {(Object.keys(LANGUAGE_CONFIGS) as Language[]).map((lang) => {
            const { label, icon: Icon, color } = LANGUAGE_CONFIGS[lang];
            return (
              <button
                key={lang}
                onClick={() => handleLanguageChange(lang)}
                className={`p-3 rounded-lg transition-all ${selectedLanguage === lang ? "bg-gray-700 shadow-lg scale-105" : "hover:bg-gray-700/50"
                  } flex flex-col items-center gap-1`}
              >
                <i className={`${Icon}`} style={{ fontSize: '27px' }} />

                <span className="text-xs">{label}</span>
              </button>
            );
          })}
        </div>

        {/* Main Content */}
        <div className="flex-1 flex">
          {/* Editor Section */}
          <div className="flex-1 border-r border-gray-700">
            <div className="h-8 bg-gray-800 border-b border-gray-700 flex items-center justify-end gap-2 px-4">
              <button
                onClick={handleFormatCode}
                className="p-1 hover:bg-gray-700 rounded transition-colors"
                title="Format Code"
              >
                <Wand2 className="w-4 h-4 text-blue-400" />
              </button>
              <button
                onClick={handleResetCode}
                className="p-1 hover:bg-gray-700 rounded transition-colors"
                title="Reset Code"
              >
                <RotateCcw className="w-4 h-4 text-yellow-400" />
              </button>
            </div>
            <div className="h-[calc(100%-2rem)]">
              <Editor
                height="100%"
                defaultLanguage={selectedLanguage}
                language={selectedLanguage}
                value={code}
                onChange={(value) => setCode(value || "")}
                theme="vs-dark"
                onMount={(editor) => setEditorInstance(editor)}
                options={{
                  minimap: { enabled: false },
                  fontSize: 14,
                  padding: { top: 16 },
                  scrollBeyondLastLine: false,
                  suggestOnTriggerCharacters: true,
                  quickSuggestions: true,
                  snippetSuggestions: "inline",
                  wordBasedSuggestions:"allDocuments",
                  parameterHints: {
                    enabled: true,
                    cycle: true,
                  },
                  suggest: {
                    showKeywords: true,
                    showSnippets: true,
                    showClasses: true,
                    showFunctions: true,
                    showVariables: true,
                    showWords: true,
                    showMethods: true,
                    showProperties: true,
                  },
                }}
              />
            </div>
          </div>

          {/* Output Section */}
          <div className="w-[40%] bg-gray-800 flex flex-col">
            <div className="h-8 bg-gray-800 border-b border-gray-700 flex items-center px-4">
              <h2 className="font-semibold">Output</h2>
            </div>
            <div className="flex-1 p-4 font-mono text-sm overflow-auto bg-gray-900 m-2 rounded-lg">
              <pre>{output}</pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
