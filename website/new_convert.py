#!/usr/bin/env python3
import sys
import subprocess
from pathlib import Path
import re

# Root folder passed as argument
if len(sys.argv) < 2:
    print("Usage: python convert_notebooks.py <root_folder>")
    sys.exit(1)

root_folder = Path(sys.argv[1])
base_dir = Path("GoFastPaced/content/docs")

# Azerbaijani chapter folders and descriptions
chapters = {
    "esaslar": "Go dilinin əsasları: dəyişənlər, sabitlər, giriş/çıxış, operatorlar və idarəetmə axını.",
    "massivler-xerite": "Massivlər, dilimlər və xəritələrlə işləmək.",
    "funksiyalar": "Funksiyalar, rekursiya, anonim funksiyalar və göstəricilər.",
    "interfeysler-generikler": "İnterfeyslər, generiklər və strukturlar.",
    "gorutinler-sehv-idare": "Gorutinlər, kanallar, səhvlərin idarəsi.",
    "modullar-qurulma": "Modullar, paketlər, qurulma və test."
}

# Notebook mapping to chapters
mapping = {
    "esaslar": [
        "Variables_DataTypesConstants.ipynb",
        "BasicInput.ipynb",
        "BasicOutput.ipynb",
        "BasicOperators.ipynb",
        "FlowControls.ipynb",
        "StringNumber.ipynb"
    ],
    "massivler-xerite": ["ArrayAndSlice.ipynb", "Map.ipynb"],
    "funksiyalar": ["FunctionsIntro.ipynb", "FunctionsRecursiveAnonymousClosure.ipynb", "Pointers.ipynb"],
    "interfeysler-generikler": ["BasicInterfaceTypeAssertions.ipynb", "GenericsTypeConstraintsGeneralInterface.ipynb", "TypeDefinitonsAliasesStructs.ipynb"],
    "gorutinler-sehv-idare": ["GoroutineChannelsSelectWaitGroup.ipynb", "Error_defer_panic_recover.ipynb"],
    "modullar-qurulma": ["ModulesPackagesBuild.ipynb", "BuildInstallTest.ipynb"]
}

def generate_tags(filename: str):
    words = re.findall(r'[A-Z][a-z]*|[a-z]+', filename)
    tags = []
    for w in words:
        if w.lower() in ("intro", "intros", "_"):
            continue
        tags.append(w)
    tags.append("Go")
    return tags

# Create chapter folders + _index.md
for i, (chapter, description) in enumerate(chapters.items(), start=1):
    chapter_dir = base_dir / chapter
    chapter_dir.mkdir(parents=True, exist_ok=True)

    index_file = chapter_dir / "_index.md"
    if not index_file.exists():
        with open(index_file, "w", encoding="utf-8") as f:
            f.write(f"""---
title: "Fəsil {i}: {chapter.replace('-', ' ').title()}"
weight: {i}
---
{description}
""")

# Convert notebooks and move them
for chapter, notebooks in mapping.items():
    chapter_dir = base_dir / chapter
    for order, nb in enumerate(notebooks, start=1):
        nb_path = root_folder / nb
        if not nb_path.exists():
            print(f"Skipping {nb} (not found in {root_folder})")
            continue

        filename = nb_path.stem
        md_output = chapter_dir / f"{filename}.md"

        # Run nbconvert with stdout
        result = subprocess.run(
            ["jupyter", "nbconvert", "--to", "markdown", "--stdout", str(nb_path)],
            capture_output=True, text=True
        )

        if result.returncode != 0:
            print(f"Error converting {nb}: {result.stderr}")
            continue

        # Generate tags
        tags = generate_tags(filename)

        # Build front matter
        title = " ".join(tags[:-1]) if len(tags) > 1 else filename.replace("_", " ")
        front_matter = f"""---
title: "{title}"
weight: {order}
tags: [{", ".join(f'"{t}"' for t in tags)}]
---
"""

        # Write final file
        with open(md_output, "w", encoding="utf-8") as f:
            f.write(front_matter + "\n" + result.stdout)

        print(f"Converted {nb} → {md_output}")
