import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Day8 {

  private String destination = "ZZZ";
  private String start = "AAA";

  private char startPart2 = 'A';
  private char destinationPart2 = 'Z';

  public static void main(String[] args) throws IOException {
    Day8 day8 = new Day8();

    List<String> allLines = Files.readAllLines(Paths.get(
        "/home/gprado/Development/private_repo/AdventCode/Day8/src/main/resources/input.txt"));

    String commands = day8.getCommands(allLines);
    List<NavigationEntry> navigationEntries = day8.parseMap(allLines);
    //Integer steps = day8.countSteps(commands, navigationEntries);
    Integer stepsPart2 = day8.countStepsFromMultipleCoordinates(commands, navigationEntries);

    //System.out.println(steps);
    System.out.println(stepsPart2);

  }

  public String[] parseMapEntry(String mapEntry) {
    return mapEntry.split(" = ");
  }

  public List<NavigationEntry> parseMap(List<String> given) {
    return given.stream().filter(g -> g.contains("=")).map(this::parseMapEntry)
        .map(this::getNavigationEntry).toList();
  }

  private NavigationEntry getNavigationEntry(String[] strings) {
    return NavigationEntryBuilder.builder().withDestination(strings[0])
        .withCoordinate(CoordinateBuilder.builder().withEntry(strings[1]).build()).build();
  }

  public String getCommands(List<String> given) {
    return given.get(0);
  }

  public Integer countSteps(String commands, List<NavigationEntry> navigationEntries) {
    boolean arrived = false;
    Integer steps = 1;
    String destination = navigationEntries.stream()
        .filter(n -> n.destination().equalsIgnoreCase(start))
        .map(n -> {
          if (commands.charAt(0) == 'R') {
            return n.coordinate().right();
          } else {
            return n.coordinate().left();
          }
        }).findFirst().get();
    while (!arrived) {

      for (int i = steps == 1 ? 1 : 0; i < commands.length(); i++) {
        destination = navigate(commands.charAt(i), navigationEntries, destination);
        steps++;
        if (destination.equalsIgnoreCase(this.destination)) {
          arrived = true;
          break;
        }
      }
    }
    return steps;
  }

  private String navigate(char direction, List<NavigationEntry> allNavigationEntries,
      String destination) {
    return allNavigationEntries.stream()
        .filter(nav -> destination.equalsIgnoreCase(nav.destination()))
        .map(n -> {
          if (direction == 'R') {
            return n.coordinate().right();
          } else {
            return n.coordinate().left();
          }
        }).findFirst().get();
  }

  public Integer countStepsFromMultipleCoordinates(String commands,
      List<NavigationEntry> navigationEntries) {
    boolean arrived = false;
    Integer steps = 1;
    List<String> destinations = navigationEntries.stream()
        .filter(n -> n.destination().charAt(2) == startPart2)
        .map(n -> {
          if (commands.charAt(0) == 'R') {
            return n.coordinate().right();
          } else {
            return n.coordinate().left();
          }
        }).collect(Collectors.toList());
    while (!arrived) {
      for (int i = steps == 1 ? 1 : 0; i < commands.length(); i++) {
        List<String> destinationsAux = new ArrayList<>();
        for (String destination : destinations) {
          destinationsAux.add(navigate(commands.charAt(i), navigationEntries, destination));
        }
        destinations = destinationsAux;
        steps++;
        if (!destinations.stream().filter(dest -> dest.charAt(2) != destinationPart2).findFirst()
            .isPresent()) {
          arrived = true;
          break;
        }
      }
    }
    return steps;
  }
}
