#!/bin/bash
dir="${1:-}"
regenerate="${2:-1}"

for author_dir in  $dir/*
do

if [ -d "$author_dir" ] && [ ! -L "$author_dir" ]; then


author=$(basename ${author_dir})
echo ${author_dir}
for x in $author_dir/*.ipynb
do
#convert to markdown
base_name=$(basename ${x})
filename="${base_name%.*}"
#jupyter nbconvert --to markdown $x  --output "${filename}${author}"  --output-dir ${PWD}
tags="tags: ["
title=""
res=$(grep -oE '[A-Z][a-z]+' <<< "$filename"  )
echo $res
for word in $res
do

  #ignore And
  if [[ $word != "And" && $word != "_" ]]
  then
    tags="$tags\"$word\", "
    echo $tags
  fi
title="$title $word"
done
git_last_date=$(git log -1  --date=short --pretty=format:'date: %cd' $x)
prepend="---\ntitle: \"$title\"\nauthor: \"${author}\"\n${git_last_date}\n$tags \"Go\" ]\n---\n"
Output_File="GoFastPaced/content/posts/${filename}_${author}.md"
if [[ $regenerate -eq 1 ]] || [[ ! -f "$Output_File" ]]
then
(echo -e $prepend; jupyter nbconvert --to markdown $x --stdout) > ${Output_File}
else
echo "skip ${Output_File}"
fi
done

fi

done