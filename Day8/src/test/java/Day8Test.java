import java.util.List;
import java.util.stream.Stream;
import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class Day8Test {

  @Test
  void test__givenMapData__parseValues() {
    var mapEntry = "AAA = (BBB, CCC)";
    Day8 day8 = new Day8();
    var map = day8.parseMapEntry(mapEntry);
    Assertions.assertThat(map).isEqualTo(new String[]{"AAA", "(BBB, CCC)"});
  }

  @Test
  void test__givenCoordinateEntry__thenReturnCoordinate() {
    var mapEntry = "AAA = (BBB, CCC)";
    Day8 day8 = new Day8();
    var map = day8.parseMapEntry(mapEntry);
    Coordinate coordinate = CoordinateBuilder.builder().withEntry(map[1]).build();
    Assertions.assertThat(coordinate.left()).isEqualTo("BBB");
    Assertions.assertThat(coordinate.right()).isEqualTo("CCC");
  }

  @Test
  void test__givenCoordinateEntry__thenReturnNavigationEntry() {
    var mapEntry = "AAA = (BBB, CCC)";
    Day8 day8 = new Day8();
    var map = day8.parseMapEntry(mapEntry);
    Coordinate coordinate = CoordinateBuilder.builder().withEntry(map[1]).build();
    NavigationEntry navigationEntry = NavigationEntryBuilder.builder().withDestination(map[0]).withCoordinate(coordinate).build();
    Assertions.assertThat(navigationEntry.destination()).isEqualTo("AAA");
    Assertions.assertThat(navigationEntry.coordinate().left()).isEqualTo("BBB");
    Assertions.assertThat(navigationEntry.coordinate().right()).isEqualTo("CCC");
  }

  public static Stream<Arguments> mapsAndSteps() {
    return Stream.of(
        Arguments.of(List.of("RL",

            "AAA = (BBB, CCC)",
            "BBB = (DDD, EEE)",
            "CCC = (ZZZ, GGG)",
            "DDD = (DDD, DDD)",
            "EEE = (EEE, EEE)",
            "GGG = (GGG, GGG)",
            "ZZZ = (ZZZ, ZZZ)"), 2),
        Arguments.of(List.of("LLR",

                "AAA = (BBB, BBB)",
            "BBB = (AAA, ZZZ)",
        "ZZZ = (ZZZ, ZZZ)"), 6)
    );
  }

  @ParameterizedTest
  @MethodSource("mapsAndSteps")
  void test__givenCommandAndCoordinates__thenCountHowManyStepsToZZZ(List<String> given, Integer expected) {
    Day8 day8 = new Day8();
    var commands = day8.getCommands(given);
    List<NavigationEntry> navigationEntries = day8.parseMap(given);
    var steps = day8.countSteps(commands, navigationEntries);
    Assertions.assertThat(steps).isEqualTo(expected);
  }

  public static Stream<Arguments> mapsAndStepsPart2() {
    return Stream.of(
        Arguments.of(List.of("LR",
            "11A = (11B, XXX)",
            "11B = (XXX, 11Z)",
            "11Z = (11B, XXX)",
            "22A = (22B, XXX)",
            "22B = (22C, 22C)",
            "22C = (22Z, 22Z)",
            "22Z = (22B, 22B)",
            "XXX = (XXX, XXX)"), 6)

    );
  }

  @ParameterizedTest
  @MethodSource("mapsAndStepsPart2")
  void test__givenCommandAndCoordinates__thenCountHowManyStepsToFromAtoZ(List<String> given, Integer expected) {
    Day8 day8 = new Day8();
    var commands = day8.getCommands(given);
    List<NavigationEntry> navigationEntries = day8.parseMap(given);
    var steps = day8.countStepsFromMultipleCoordinates(commands, navigationEntries);
    Assertions.assertThat(steps).isEqualTo(expected);
  }
}
