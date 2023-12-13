import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Day13 {

  public static void main(String[] args) throws IOException {
    Day13 day13 = new Day13();

    List<String> allLines = Files.readAllLines(Paths.get(
        "/home/gprado/Development/private_repo/AdventCode/Day13/src/main/resources/input.txt"));

    List<String[]> patterns = day13.parsePatterns(allLines);

    int summary = 0;
    for (int i = 0; i < patterns.size(); i++) {
      var mirrorsFound = day13.findMirroredPatternWithSmudge(patterns.get(i));
      summary += mirrorsFound.getSummary();

    }

    System.out.println(summary);

  }

  public Mirror findMirroredPattern(String[] pattern) {
    for (int i = 0; i < pattern.length - 1; i++) {
      if(pattern[i].equalsIgnoreCase(pattern[i+1])) {
        int counter = i+1;
        boolean isMirrored = true;
        for (int j = i-1; j >= 0 && counter + 1< pattern.length; j--) {
          counter++;
          if(!pattern[j].equalsIgnoreCase(pattern[counter]) ) {
            isMirrored = false;
          }
        }
        if (isMirrored) {

          return new Mirror(new int[]{i+1, i+2}, Orientation.HORIZONTAL);

        }
      }
    }

    Map<Integer, String> columns = new HashMap<>();
    for (int j = 0; j < pattern[0].length() - 1; j++) {
      String column1 = "";
      String column2 = "";
      for (int i = 0; i < pattern.length; i++) {
        column1 += pattern[i].charAt(j);
        column2 += pattern[i].charAt(j+1);
      }
      columns.put(j, column1);
      columns.put(j+1, column2);
      if(column1.equalsIgnoreCase(column2)) {
        var isMirrored = true;
        int k = j+1;
        int j2 = j;
        while(k < pattern[0].length() - 1 && j2 > 0) {
          if (!columns.get(--j2).equalsIgnoreCase(getColumnsValue(++k, pattern))) {
            isMirrored = false;
          }
        }
        if (isMirrored) {
          return new Mirror(new int[]{j+1, j+2}, Orientation.VERTICAL);
        }
      }
    }

    return null;
  }

  public MirrorCheck checkMirror(String s1, String s2) {
    if (s1.equalsIgnoreCase(s2)) {
      return new MirrorCheck(true, false);
    } else {
      int difference = 0;
      for (int i = 0; i < s1.length(); i++) {
        if (s1.charAt(i) != s2.charAt(i)) {
          difference++;
        }
      }
      if (difference == 1) {
        return new MirrorCheck(true, true);
      }
    }
    return new MirrorCheck(false, false);
  }
  public Mirror findMirroredPatternWithSmudge(String[] pattern) {
    for (int i = 0; i < pattern.length - 1; i++) {
      MirrorCheck mirrorCheck = checkMirror(pattern[i], pattern[i + 1]);
      if(mirrorCheck.isMirrored()) {
        boolean isSmudged = mirrorCheck.isSmudged();
        int counter = i+1;
        boolean isMirrored = true;
        MirrorCheck mirrorCheckSeq = null;
        for (int j = i-1; j >= 0 && counter + 1< pattern.length; j--) {
          counter++;
          mirrorCheckSeq = checkMirror(pattern[j], pattern[counter]);
          if(!mirrorCheckSeq.isMirrored()) {
            isMirrored = false;
          } else {
            isSmudged = isSmudged || mirrorCheckSeq.isSmudged();
          }
        }
        if (isMirrored && isSmudged) {

          return new Mirror(new int[]{i+1, i+2}, Orientation.HORIZONTAL);

        }
      }
    }

    Map<Integer, String> columns = new HashMap<>();
    for (int j = 0; j < pattern[0].length() - 1; j++) {
      String column1 = "";
      String column2 = "";
      for (int i = 0; i < pattern.length; i++) {
        column1 += pattern[i].charAt(j);
        column2 += pattern[i].charAt(j+1);
      }
      columns.put(j, column1);
      columns.put(j+1, column2);
      MirrorCheck mirrorCheck = checkMirror(column1, column2);
      if(mirrorCheck.isMirrored()) {
        var isMirrored = true;
        boolean isSmudged = mirrorCheck.isSmudged();
        int k = j+1;
        int j2 = j;
        MirrorCheck mirrorCheckSeq = null;
        while(k < pattern[0].length() - 1 && j2 > 0) {
          mirrorCheckSeq = checkMirror(columns.get(--j2), getColumnsValue(++k, pattern));
          if(!mirrorCheckSeq.isMirrored()) {
            isMirrored = false;
          } else {
            isSmudged = isSmudged || mirrorCheckSeq.isSmudged();
          }
        }
        if (isMirrored && isSmudged) {
          return new Mirror(new int[]{j+1, j+2}, Orientation.VERTICAL);
        }
      }
    }

    return null;
  }

  private String getColumnsValue(int index, String[] pattern) {
    String column = "";
    for (int i = 0; i < pattern.length; i++) {
      column += pattern[i].charAt(index);
    }
    return column;
  }

  public List<String[]> parsePatterns(List<String> fullAnnotations) {
    List<String[]> patterns = new ArrayList<>();
    int counter = 0;
    for (int i = 0; i < fullAnnotations.size(); i++) {
      if(fullAnnotations.get(i).isEmpty()) {
        patterns.add(fullAnnotations.subList(i-counter, i).toArray(new String[0]));
        counter = 0;
      } else {
        counter += 1;
      }
    }
    patterns.add(fullAnnotations.subList(fullAnnotations.size()-counter, fullAnnotations.size()).toArray(new String[0]));
    return patterns;
  }
}
