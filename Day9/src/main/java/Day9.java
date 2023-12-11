import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Day9 {

  public static void main(String[] args) throws IOException {
    Day9 day9 = new Day9();

    List<String> allLines = Files.readAllLines(Paths.get(
        "/home/gprado/Development/private_repo/AdventCode/Day9/src/main/resources/input.txt"));

    Integer i = day9.calculateAllEntries(allLines);
    List<Oasis> allOasis = day9.buildAllOasis(allLines);
    System.out.println(i);


  }

  public Integer calculateAllEntries(List<String> allLines) {
    Integer total = 0;
    for (String entry : allLines) {
      Oasis oasis = new Oasis();
      Integer[] values = Arrays.stream(entry.split(" ")).map(Integer::parseInt)
          .toArray(Integer[]::new);
      oasis.buildHistory(values);
      oasis.extrapolateToZeroes();
      total += oasis.calculateNextValueReversed();
    }
    return total;
  }

  public List<Oasis> buildAllOasis(List<String> entries) {
    List<Oasis> allOasis = new ArrayList<>();
    for (String entry : entries) {
      Oasis oasis = new Oasis();
      Integer[] values = Arrays.stream(entry.split(" ")).map(Integer::parseInt)
          .toArray(Integer[]::new);
      oasis.buildHistory(values);
      allOasis.add(oasis);
    }
    return allOasis;
  }
}