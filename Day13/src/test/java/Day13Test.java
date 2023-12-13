import java.util.Arrays;
import java.util.stream.Stream;
import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class Day13Test {

  public static Stream<Arguments> annotations() {
    var pattern = new String[]{"#.##..##.",
        "..#.##.#.",
        "##......#",
        "##......#",
        "..#.##.#.",
        "..##..##.",
        "#.#.##.#."};

    var pattern2 = new String[]{"#...##..#",
        "#....#..#",
        "..##..###",
        "#####.##.",
        "#####.##.",
        "..##..###",
        "#....#..#"};

    return Stream.of(
        Arguments.of(pattern, new Mirror(new int[]{5, 6}, Orientation.VERTICAL)),
        Arguments.of(pattern2, new Mirror(new int[]{4, 5}, Orientation.HORIZONTAL))

    );
  }

  @ParameterizedTest
  @MethodSource("annotations")
  void givenPatternAnnotation__findMirroredPattern(String[] pattern, Mirror mirrors) {
    Day13 day13 = new Day13();
    var mirrorsFound = day13.findMirroredPattern(pattern);
    Assertions.assertThat(mirrorsFound).isEqualTo(mirrors);
  }

  @Test
  void givenPatternAnnotation__findMirroredPatternSummaryCount() {
    var pattern = new String[]{
        "#.##..##.",
        "..#.##.#.",
        "##......#",
        "##......#",
        "..#.##.#.",
        "..##..##.",
        "#.#.##.#.",
        "",
        "#...##..#",
        "#....#..#",
        "..##..###",
        "#####.##.",
        "#####.##.",
        "..##..###",
        "#....#..#",
        "",
        "..##.##.##.##.##.",
        "####..##..##..###",
        "..##..######..##.",
        "..#.##.#..#.##.#.",
        "##..###.##..##..#",
        "##..##########..#",
        "##.#....##....#.#",
        "####.#.#..#.#.###",
        "....####..####...",
        "..##..#....#..##.",
        ".......####......",
        "....#.#....#.#...",
        "####..........###",
        "..###.##..##.###.",
        "###....####....##",
        "####....##....###",
        "###..#.####.#..##"
    };
    Day13 day13 = new Day13();
    var patterns = day13.parsePatterns(Arrays.asList(pattern));
    int summary = 0;
    for (int i = 0; i < patterns.size(); i++) {
      var mirrorsFound = day13.findMirroredPattern(patterns.get(i));
      summary += mirrorsFound.getSummary();

    }

    Assertions.assertThat(summary).isEqualTo(406);
  }

  @Test
  void givenPatternAnnotation__findMirroredPatternSummaryCountWithSmudge() {
    var pattern = new String[]{
        "#.##..##.",
        "..#.##.#.",
        "##......#",
        "##......#",
        "..#.##.#.",
        "..##..##.",
        "#.#.##.#.",
        "",
        "#...##..#",
        "#....#..#",
        "..##..###",
        "#####.##.",
        "#####.##.",
        "..##..###",
        "#....#..#"
    };
    Day13 day13 = new Day13();
    var patterns = day13.parsePatterns(Arrays.asList(pattern));
    int summary = 0;
    for (int i = 0; i < patterns.size(); i++) {
      var mirrorsFound = day13.findMirroredPatternWithSmudge(patterns.get(i));
      summary += mirrorsFound.getSummary();

    }

    Assertions.assertThat(summary).isEqualTo(400);
  }
}